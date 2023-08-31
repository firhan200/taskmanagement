import axios from "axios";
import { useState } from "react";
import Button from "src/atoms/Button/Button";
import Loading from "src/atoms/Loading/Loading";
import useAuth from "src/hooks/useAuth";

const apiUrl = import.meta.env.VITE_API_URL

export default function GeneratePage() {
    const { token } = useAuth()
    const [loading, setLoading] = useState(false)

    const starGenerate = async () => {
        setLoading(true)

        try{
            const res = await axios(`${apiUrl}generate`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            const data = res.data
            console.log(data)
        }catch(err){
            console.log(err)
        }

        setLoading(false)
    }

    const confirmation = () => {
        const confirm = window.confirm("Generate Random Data?")
        if (confirm) {
            starGenerate()
        }
    }

    return (
        <div className="my-20">
            <div className="text-center">
                {
                    loading ? (
                        <Loading />
                    ) : (
                        <Button onClick={() => confirmation()} colorType="danger" size="lg">Generate</Button>
                    )
                }
            </div>
        </div>
    );
}