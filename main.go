package main

import (
	// "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/spf13/viper"
    geometry "github.com/andrewlunde/thetaoffchaingo"

    "github.com/andrewlunde/greetings"

    "github.com/thetatoken/theta/query"

)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/offchain/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/offchain/links" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    // data := []byte("V1 of student's called")
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(200)
    fmt.Fprintf(w, "<a href=\"/offchain/links\">links</a><br />\n")
    fmt.Fprintf(w, "<a href=\"/offchain/hello\">hello</a><br />\n")
    fmt.Fprintf(w, "<a href=\"/index.html\">index</a><br />\n")
    fmt.Fprintf(w, "<a href=\"/form.html\">form</a><br />\n")
    // w.Write(data)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

// Alice=2E833968E5bB786Ae419c4d13189fB081Cc43bab
// Bob=70f587259738cB626A1720Af7038B8DcDb6a42a0
// Carol=cd56123D0c5D6C1Ba4D39367b88cba61D93F5405
// rid=rid1000001

//cmd='thetacli tx reserve --chain="privatenet" --async --from='$Alice' --fund='$rfund' --collateral='$rcoll' --duration='$rfdurationblocks' --resource_ids='$rid' --seq='$ans' --password=qwertyuiop'


func main() {

	ellipse := geometry.Ellipse{
		9, 2,
	}
	fmt.Println(ellipse.GetEccentricity())

    // A slice of names.
    names := []string{"Andrew", "Dana", "Birdie"}

    // Request a greeting message.
    message, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)

    endpoint := "https://theta-dev-theta-privatenet.cfapps.eu10.hana.ondemand.com/rpc"
    address := "0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"
    query.Account(endpoint,address)


    fileServer := http.FileServer(http.Dir("./static")) 
    http.Handle("/", fileServer) 

    http.HandleFunc("/offchain/hello", helloHandler)

    http.HandleFunc("/offchain/links", linksHandler)

    http.HandleFunc("/offchain/form", formHandler)

    viper.SetDefault("port", "8080")
    viper.AutomaticEnv()
    port := viper.GetString("port")
    fmt.Printf("Starting server at port %s\n", port)
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}

