import { defineStore } from "pinia";
import { ref } from "vue";
import { Modality } from "../enums/modality";
import type { LocalStorageOperationResult } from "../types/LocalStorage";
import type { User } from "../types/User";
import { Seniority } from "../enums/seniority";
import { Category as CategoryEnum } from "../enums/category";


export const useUserStore = defineStore('user', () =>{
  const description = ref<string>('')
  const level = ref<Seniority>(Seniority.ANY)
  const skills = ref<string[]>([])
  const location = ref<string>('')
  const modality = ref<Modality>(Modality.ANY)
  const minSalary = ref<number>(0)
  const maxSalary = ref<number>(0)
  const category = ref<CategoryEnum>(CategoryEnum.NotCategory)
  const hostSelected = ref<number[]>([0,1])

  const resetData = () => {
    description.value = ''
    level.value = Seniority.JUNIOR
    skills.value = []
    location.value = ''
    modality.value = Modality.ANY
    minSalary.value = 0
    maxSalary.value = 0
    category.value = CategoryEnum.NotCategory
    hostSelected.value = [0,1]
  }

  const addSkill = (skill:string) =>{
    for(const sk of skills.value){
      if(sk.toLowerCase() === skill.toLowerCase()){
        return
      }
    }

    skills.value =  [...skills.value,skill]
  }

  const removeSkill = (skill: string) => {
    skills.value = skills.value.filter(sk => sk !== skill)
  }

  const clearSkills = () => {
    skills.value = []
  }

  const loadDataFromLocalStorage = ():LocalStorageOperationResult =>  {
    try {
      const user = localStorage.getItem('user')
      if (user) {
        const userData:User | null = JSON.parse(user)
        if (userData) {
          description.value = userData.description || ''
          level.value = userData.level as Seniority || Seniority.ANY
          skills.value = userData.skills || []
          location.value = userData.location || ''
          modality.value = userData.modality as Modality || Modality.ANY
          minSalary.value = userData.minSalary || 0
          maxSalary.value = userData.maxSalary || 0
          category.value = userData.category as CategoryEnum || CategoryEnum.NotCategory
          hostSelected.value = userData.hostSelected as number[] || [0,1]
        }
      }
      return { ok: true, error: '' }
    } catch (error) {
      console.error('Error loading data from localStorage:', error)
      return { ok: false, error: 'Error loading the data' }
    }
  }
  
  const saveDateInLocalStorage = ():LocalStorageOperationResult => {
    try {
      localStorage.setItem('user', JSON.stringify({
        description: description.value,
        level: level.value,
        skills: skills.value,
        location: location.value,
        modality: modality.value,
        minSalary: minSalary.value,
        maxSalary: maxSalary.value,
        category: category.value,
        hostSelected: hostSelected.value,
      }))

      return { ok: true, error: '' }
    } catch (error) {
      console.error('Error saving data in localStorage:', error)
      return { ok: false, error: 'Error saving data' }
    }

  }

  return {
    description,
    level,
    skills,
    location,
    modality,
    minSalary,
    maxSalary,
    category,
    hostSelected,
    resetData,
    addSkill,
    removeSkill,
    clearSkills,
    loadDataFromLocalStorage,
    saveDateInLocalStorage
  }

})