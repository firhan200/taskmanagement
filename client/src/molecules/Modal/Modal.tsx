import useModal from "src/hooks/useModal";

export default function Modal() {
    const { isShow } = useModal();

    return (
        <>
            <input type="checkbox" checked={isShow} id="notification_modal" className="modal-toggle" />
            <div className="modal">
                <div className="modal-box">
                    <h3 className="font-bold text-lg">Hello!</h3>
                    <p className="py-4">This modal works with a hidden checkbox!</p>
                    <div className="modal-action">
                        <label htmlFor="notification_modal" className="btn">Close!</label>
                    </div>
                </div>
            </div>
        </>
    );
}