package main

import (
	"context"
	"fmt"
	"log"
	"social_network/constant/noti"
	"social_network/dto"
	"social_network/service"

	"github.com/joho/godotenv"
)

func main() {
	//cmd.Execute()
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
	}

	service, _ := service.GenerateUserService()
	res, _ := service.UpdateUser(dto.UpdateUserReq{
		FullName: "Hi hi hi",
		UserId:   "00a51055-949c-4421-9210-337530f38a7e",
	}, "00a51055-949c-4421-9210-337530f38a7e", context.Background())

	log.Println("Res: ", res)
}

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }

// service, _ := service.GenerateUserService()

// res1, res2, err := service.Login(dto.LoginRequest{
// 	Email:    "john.doe@example.com",
// 	Password: "@A123456a78",
// }, context.Background())

// log.Println("Res 1: " + res1)
// log.Println("Res 2: " + res2)
// log.Println("Error: " + err.Error())

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }

// service, _ := service.GenerateUserService()

// res, err := service.CreateUser(dto.CreateUserReq{
// 	Username:      "janesmith",
// 	FullName:      "Jane Smith",
// 	Email:         "jane@example.com",
// 	Password:      "@A12345678a",
// 	DateOfBirth:   time.Now(),
// 	ProfileAvatar: "avatar_url.jpg",
// }, "", context.Background())

// log.Println("Res: " + res)
// log.Println("Error: " + err.Error())

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }
// service, _ := service.GeneratePostService()

// res := service.GetPosts(context.Background())
// log.Println("Res: ", res)

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }

// service, _ := service.GenerateLikeService()
// res := service.DoLike(dto.DoLikeReq{
// 	ActorId:    "00a51055-949c-4421-9210-337530f38a7e",
// 	ObjectId:   "p2",
// 	ObjectType: "post",
// }, context.Background())

// log.Println("Result: ", res.Error())

// if err := godotenv.Load(); err != nil {
// 	log.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
// }

// service, _ := service.GenerateCommentService()
// res := service.EditComment(dto.EditCommentRequest{
// 	ActorId:   "00a51055-949c-4421-9210-337530f38a7e",
// 	CommentId: "242061f1-c255-47af-a13a-d462056fa8f3",
// 	Content:   "comment 1.1",
// }, context.Background())

// log.Println("Res: ", res.Error())
