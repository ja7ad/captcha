package main

import (
	"context"
	"fmt"
	"github.com/ja7ad/captcha/_example/protobuf/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := echo.NewEchoServiceClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"x-challenge-key": "0.MMdulva8VNcSGNXhY5VCqBNE4HbJ7weG3EYN0L3bGy9o1Muj0rxe9-Ir4WWV8sEQ3vu1cXhW6JK_6F-B2bXGaPSHuROG3P6e6z7UmD6u09xii4ULh3wEx1912kOFFvv_urh4whHmvYQgy1e4kYlG7ayCN5v6uMba_nADxjZttvRB4-zV6_Rbrke0K1LERfWvnWIe0xrYWg6_2r06UDQQ3PQ5rEtXb9O8hvuKSeeY4YsSew4Sv4uIi3yfVLT9HXMVhH1ZLsE0f35FGF4uXyxOb1zVIq74Kyh2H0FmmITLFkpGcLW88W4OZvzC55FzMNGSTGWngmlA9Owpiq7lmR3l1TgeC00lRuIFrnh-ea3vZTNixT2Sp8jzv3r5GWkAz_elwI65Aa36E04RhSyte9y7AWewSMFrUGCp9dE9GIZFhqr4xlrlzyELohl_G4LGBfH0__xKjzThzCXsdPjFkJgzcJo1Twv2g47wgQ4Aee4rg0AnbdMlOl__8l7NHQdb_etniVP7852AH7GmRhrObDdDpRjL6wG-y4szdj_EU49cSzlx5Hu7OnXRMSiGNMCwPsQrTKf13DQGip21cMxP67OYQ0w7Q3NK_R5WPUDiajTBKLgZUdAC6v8Oa8N2z_8-tyucnCS30P_xb0Bm0fdZ5bEOSfde45qOBaQ2W-27ycJI2hB4vpgWP2al23F2NKyk4gwP-0y2EF2EhsNGrraoxKlMgwVPjzvkC4NNiEzigugUSD1P65qx6TMsRk8_oNhUxeoDy3REdYtj_zCA4PLdsC7v3_h__KF5QQciVrq-Q8AP12HCfVpac-PvDduhOzX5hXZwX5kCbzzA3o3nNzeevRVWJA.mt5NFc6QcAW046f5-HMoXw.de287b261f0a3b893e83b779d76d45b3d0699db530dc866e022895d5a67aa34b",
	}))

	resp, err := client.Echo(ctx, &echo.EchoRequest{
		Name: "foobar",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Message)
}
