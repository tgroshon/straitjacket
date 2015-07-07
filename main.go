package main

// sudo ./install_apparmor_profiles
// docker build -t cstest languages/csharp
// docker build -t straitjacket . && docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v /tmp:/tmp --rm -p 8081:8081 straitjacket

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	StartServer(":8081")
}
