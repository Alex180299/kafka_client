function Produce() {
    function sendMessage() {
        console.log("send message")
    }

    return (
        <div>
            <h5>Topic</h5>
            <input type="text" placeholder="example.topic" id="topic" />

            <h5>Key</h5>
            <input type="text" placeholder="example.key" id="key" />

            <h5>Value</h5>
            <input type="text" placeholder="example.value" id="value" />

            <button onClick={sendMessage}>Send</button>
        </div>
    );
}

export default Produce;