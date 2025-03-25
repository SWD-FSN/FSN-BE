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
