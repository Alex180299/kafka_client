import React, { useState } from 'react';

function ListenMessages() {
    const [messages, setMessages] = useState([])

    msgs = []
    for (let i = 0; i < messages.length; i++) {
        this.msgs.push(<p>{messages[i]}</p>)
    }

    return (
        <div>
            {msgs}
        </div>
    );
}

export default ListenMessages;