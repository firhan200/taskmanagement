const FormControlWrapper = ({ title, children }: { title: string, children: React.ReactNode }) => {
    return (
        <div className="form-control">
            <label className="label">
                <span className="label-text">{ title }</span>
            </label>
            
            {
                children
            }
        </div>
    );
}

export default FormControlWrapper;