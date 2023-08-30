import { create } from "zustand"

type ModalStoreState = {
    isShow: boolean
    show: (show: boolean) => void
}

const modalStore = create<ModalStoreState>()((set) => ({
    isShow: false,
    show: (show: boolean) => set(() => {
        return {
            isShow: show
        }
    })
}))

export default function useModal(){
    const [ isShow, show ] = modalStore((state) => [state.isShow, state.show])
    return {
        isShow,
        show
    }
}