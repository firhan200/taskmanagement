import { create } from "zustand"

type ModalStoreState = {
    message: string | null
    cb: null | (() => void)
    show: (message: string, cb: null | (() => void)) => void
    hide: () => void
}

export const modalStore = create<ModalStoreState>()((set) => ({
    message: null,
    cb: null,
    show: (message: string, cb: null | (() => void)) => set(() => {
        return {
            message: message,
            cb: cb
        }
    }),
    hide: () => set(() => ({
        message: null,
        cb: null
    }))
}))

export default function useModal(){
    const [ message, cb, show, hide ] = modalStore((state) => [state.message, state.cb, state.show, state.hide])
    return {
        message,
        cb,
        show,
        hide
    }
}