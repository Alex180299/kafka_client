import Address from '../address/Address'
import Produce from '../produce/Produce'
import Listen from '../listen/Listen'

function KafkaClient() {
    return (
        <div className="KafkaClient">
            <Address />
            <Produce />
            <Listen />
        </div>
    );
}

export default KafkaClient;