package main

import "social_network/cmd"

func main() {
	cmd.Execute()
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
