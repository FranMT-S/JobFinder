import { defineStore } from "pinia"
import { onMounted, ref, watch } from "vue"
import type { Category } from "../types/Categories"
import { fetchSkills } from "../services/fetchSkills"
import { fetchCategories } from "../services/fetchCategories"
import type { HostScrapper } from "../types/HostScrapper"
import { fetchHost } from "../services/fetchHost"

export const useGlobalStore = defineStore('global', () => {
    

    const isDarkMode = ref(localStorage.getItem('darkMode') === 'true')
    const isSalaryRange = ref(localStorage.getItem('salaryRange') === 'true')
    const maxPage = ref<number>(Number(localStorage.getItem('maxPage')) || 1)
    const theme = ref(isDarkMode.value ? 'dark' : 'light')
    const skillsList = ref<string[]>([])
    const categoriesList = ref<Category[]>()
    const hostList = ref<HostScrapper[]>()


    const initAllData = async () => {
      await loadSkills()
      await loadCategories()
      await loadHost()
      updateDarkMode()
    }

    watch(isDarkMode, () => {
        try {
          localStorage.setItem('darkMode', isDarkMode.value.toString())
        } catch (error) {
          console.error('Error saving data in localStorage:', error)
        }
    })

    watch(isSalaryRange, () => {
      try {
        localStorage.setItem('salaryRange', isSalaryRange.value.toString())
      } catch (error) {
        console.error('Error saving data in localStorage:', error)
      }
    })

    watch(isDarkMode, () => {
      theme.value = isDarkMode.value ? 'dark' : 'light'
      updateDarkMode()
    })

    watch(maxPage, () => {
      try {
        localStorage.setItem('maxPage', maxPage.value.toString())
      } catch (error) {
        console.error('Error saving data in localStorage:', error)
      }
    })

    const updateDarkMode = () => {
      if(isDarkMode.value){
        document.body.classList.add('dark')
      }else{
        document.body.classList.remove('dark')
      }
    }

    const toggleDarkMode = () => {
        isDarkMode.value = !isDarkMode.value
    }

    const loadSkills = async () => {
        const [err,result] = await fetchSkills()

        if(err){
          console.error(err)
          return
        }

        skillsList.value = result || []
    }

    const loadCategories = async () => {
        const [err,result] = await fetchCategories()

        if(err){
            console.error(err)
            return
        }

        categoriesList.value = result || []
    }

    const loadHost = async () => {
      const [err,result] = await fetchHost()

      if(err){
        console.error(err)
        return
      }

      hostList.value = result || []
    }

    return {
        isDarkMode,
        isSalaryRange,
        maxPage,
        theme,
        toggleDarkMode,
        loadSkills,
        loadCategories,
        loadHost,
        initAllData,
        skillsList,
        categoriesList,
        hostList
    }
})
    
