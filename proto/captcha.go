package proto

import (
	"context"
	"errors"
	"fmt"
	"github.com/ja7ad/captcha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

const _defaultHeaderKey = "x-challenge-key"

type ProtoCaptcha struct {
	cap             captcha.Captcha
	customHeaderKey string
}

// NewCaptcha create captcha object for access to middleware method
func NewCaptcha(provider captcha.Provider, secretKey, customHeaderKey string) (*ProtoCaptcha, error) {
	c, err := captcha.New(provider, secretKey)
	if err != nil {
		return nil, err
	}

	if customHeaderKey == "" {
		customHeaderKey = _defaultHeaderKey
	}

	return &ProtoCaptcha{
		cap:             c,
		customHeaderKey: customHeaderKey,
	}, nil
}

// UnaryServerInterceptor captcha validator for unary server interceptor
func (c *ProtoCaptcha) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		cp, err := c.extractCaptchaOptionFromDescriptor(strings.Replace(info.FullMethod[1:], "/", ".", -1))
		if err != nil {
			return nil, err
		}

		challengeKey, err := c.challengeKeyFromHeader(ctx)
		if err != nil {
			return nil, err
		}

		if cp.CheckChallenge {
			valid, err := c.cap.Validate(ctx, challengeKey, "")
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

			if !valid {
				return nil, status.Error(codes.InvalidArgument, "invalid captcha challenge key")
			}
		}

		return handler(ctx, req)
	}
}

// StreamServerInterceptor captcha validator for stream server interceptor
func (c *ProtoCaptcha) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		cp, err := c.extractCaptchaOptionFromDescriptor(strings.Replace(info.FullMethod[1:], "/", ".", -1))
		if err != nil {
			return err
		}

		challengeKey, err := c.challengeKeyFromHeader(stream.Context())
		if err != nil {
			return err
		}

		if cp.CheckChallenge {
			valid, err := c.cap.Validate(stream.Context(), challengeKey, "")
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			if !valid {
				return status.Error(codes.InvalidArgument, "invalid captcha challenge key")
			}
		}

		return handler(srv, stream)
	}
}

func (c *ProtoCaptcha) extractCaptchaOptionFromDescriptor(methodName string) (*Captcha, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(methodName))
	if err != nil {
		return nil, err
	}

	methodDesc, ok := desc.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, errors.New("captcha: failed to assertion method descriptor")
	}

	options, ok := methodDesc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, errors.New("captcha: failed to assertion descriptor options")
	}

	capExt := proto.GetExtension(options, E_Captcha)

	cp, ok := capExt.(*Captcha)
	if !ok {
		return nil, errors.New("captcha: failed to assertion captcha object with proto extension")
	}

	return cp, nil
}

func (c *ProtoCaptcha) challengeKeyFromHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("captcha: no metadata found")
	}

	vals := md.Get(c.customHeaderKey)
	if len(vals) == 0 {
		return "", fmt.Errorf("no header %s found", c.customHeaderKey)
	}

	return vals[0], nil
}
