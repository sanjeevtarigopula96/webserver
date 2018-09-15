package main

import  (
        "net/http"
        "os"
        "time"
        "encoding/json"
)

type resource struct  {
        Hostname string `json:"hostname"`
        CurrentTime string `json:"servertime"`
}

func main()  {
        http.HandleFunc("/getresource", getresource);
        http.ListenAndServe(":8000", nil);
}

func getresource(w http.ResponseWriter, r * http.Request)  {
        hostname, _ := os.Hostname();
        data := resource{Hostname:hostname, CurrentTime:time.Now().String()}
        json.NewEncoder(w).Encode(data);
}


