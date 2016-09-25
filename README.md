# Reservoir sampling as a service

I chose Go to solve the challenge stated in the pdf ps.pdf

My solution splits algorithm R into two parts and doesn't use the filling part since the array is copied over previously

It allows reservoir sampling as a service with providing the following require features: 
- Time complexity of a displacement is O(1)
- Uses O(k) space complexity
- Supports multiple sessions

You can run the service using `go run final.go`. It runs on `localhost:8080`

The commands are: 
- Starting a session:
    - `curl --request POST http://127.0.0.1:8080/start/{session_name}/{list of numbers} -v`
        
        Example: `curl --request POST http://127.0.0.1:8080/start/awake/"1,2,3,4,5" -v`
- Uploading a new number:
    - `curl --request POST http://127.0.0.1:8080/displace/{session_name}/{number to try} -v`

        Example: `curl --request POST http://127.0.0.1:8080/start/awake/9 -v`
- Closing a session:
    - `curl --request GET http://127.0.0.1:8080/close/{session_name} -v`

        Example: `curl --request GET http://127.0.0.1:8080/close/awake -v`

Looking forward to your comments. Regardless of next steps, I'd appreciate if you could give me your feedback on this work and tips as to how I could be better. 

You may contact me at avivaid2018@u.northwestern.edu or comment on this repo
