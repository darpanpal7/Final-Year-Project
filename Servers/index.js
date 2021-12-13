const express = require("express")
const app = express()

const PORT = process.argv[2]
const LIM = 100

const getRandomInt = () => {
    return Math.floor(Math.random() * LIM)
}

app.get("/:id", (req, res) => {
    console.log("Get Request on Server : " + PORT)

    num = req.params["id"]
    
    var val = 0; 
    for(var i = 0; i < num; i++) 
        val += getRandomInt()
    
    result = String(val)
    res.send(result)
})

app.get('/', (req,  res) => {
    res.send("Hello World")
    console.log("Get Request on Server : " + PORT)
})

app.listen(PORT, () => {
    console.log("Server running on port : " + PORT)
})