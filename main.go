package main 

import ( 
    "log" 
    "net/http" 
) 

func main() { 
    // Set up the file server to serve files from the "public" directory
    fs := http.FileServer(http.Dir("public")) 

    // Use the StripPrefix method to serve index.html as the default page at the root URL path
    http.Handle("/", http.StripPrefix("/", fs)) 

    // Start the server and listen on port 80
    log.Println("Listening on port 80...")
    err := http.ListenAndServe(":80", nil) 
    if err != nil { 
        log.Fatal(err) 
    } 
}