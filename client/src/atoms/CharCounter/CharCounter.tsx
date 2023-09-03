export default function CharCounter({ current, max }: { current: number, max: number }) {
    return (
        <div className="text-end">
            { `${current}/${max}` }
        </div>
    );
}