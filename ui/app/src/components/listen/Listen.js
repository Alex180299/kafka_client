import React, { useState } from 'react';

function Listen() {
    const [topic, setTopic] = useState("");

    function listen() {
        fetch("/listen", {
            method: "POST",
            body: {
                address: "kafka",
                topic: "topic"
            }
        }).then(res => {
            if(res.ok) {
                console.log(res.text())
            } else {
                console.log("error: " + res.text())
            }
        })
    }

    function handleChangeTopic(e) {
        setTopic(e.target.value)
      };

    return (
        <div>
            <h5>Topic</h5>
            <input type="text" placeholder="example.topic" id="topic" value={topic} onChange={handleChangeTopic} />

            <button onClick={listen}>Listen</button>
        </div>
    );
}

export default Listen;