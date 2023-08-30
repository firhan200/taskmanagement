import Button from "src/atoms/Button/Button";
import Typography from "src/atoms/Typography/Typography";
import useModal from "src/hooks/useModal";

export default function Modal() {
    const { message, cb, hide } = useModal()

    const btnActionClick = () => {
        if(cb !== null){
            cb()
        }

        hide()
    }

    return (
        <dialog open={message !== null} className="modal">
            <div className="modal-box">
                <Typography size="md">{ message }</Typography>
                <div className="modal-action">
                    <Button colorType="primary" size="md" onClick={() => btnActionClick()}>OK</Button>
                </div>
            </div>
        </dialog>
    );
}