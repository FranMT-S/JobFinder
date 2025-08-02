import { ref } from "vue"

export const useToast = () => {
  const message = ref('')
  const show = ref(false)

  const trigger = (msg: string,duration:number = 3000) => {
    message.value = ''
    setTimeout(() => {
      message.value = msg
      show.value = true
      setTimeout(() => show.value = false, duration)
    }, 0);
  }

  return { message, show, trigger }
}