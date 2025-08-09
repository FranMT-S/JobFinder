
<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref, watch } from 'vue'
import { useGlobalStore } from '../stores/globalStore'
import { useUserStore } from '../stores/userStore'
import ChipList from './ChipList.vue'
import ModalitySelect from './ModalitySelect.vue'
import SenioritySelect from './SenioritySelect.vue'
import SkillsSelect from './SkillsSelect.vue'


const { description, level, skills, location, modality, minSalary, maxSalary,category,hostSelected } = storeToRefs(useUserStore())
const { saveDateInLocalStorage, loadDataFromLocalStorage, removeSkill, clearSkills } = useUserStore()


const emit = defineEmits(['onError'])
const {categoriesList, isSalaryRange} = storeToRefs(useGlobalStore())
const specifySalary = ref(false)

loadDataFromLocalStorage()



let timer: number;

watch(isSalaryRange, () => {
  if(!specifySalary.value){
    minSalary.value = 0
    maxSalary.value = 0
  }
})

watch([
  description,
  level,
  skills,
  location,
  modality,
  minSalary,
  maxSalary,
  category,
  hostSelected
], () => {
  
  clearTimeout(timer)
  timer = setTimeout(() => {
    saveDateInLocalStorage()
  }, 400)
})

const onClickChip = (index: number) => {
  const skill = skills.value[index]
  removeSkill(skill)
}


</script> 

<template>
  
  <form class="max-w-2xl  p-6 bg-black_primary h-full  overflow-auto  dark:text-white" autocomplete="off">
 
    <SenioritySelect :seniority="level" @update:seniority="level = $event" />
    <v-select
      v-model="category"
       density="compact"
      color="primary"
      :items="categoriesList"
      item-title="name"
      item-value="id"
      label="Category"
      @update:modelValue="category = $event"
    />
    <ModalitySelect class="mt-4" :modality="modality"  />

    <div class="flex-col gap-4">
      <div class="flex items-center gap-2 item-center"> 
        <v-switch
          v-model="isSalaryRange"
          color="primary"
          density="compact"
          prepend-icon="mdi-currency-usd-off"
          :title="!isSalaryRange ? 'Any range' : 'Salary range'"
        >
          <template #label>
            <v-icon class="cursor-default" icon="mdi-currency-usd" />
          </template>
        </v-switch>
      </div>
      <div class="flex  gap-4" v-if="isSalaryRange">
        <v-number-input
          color="primary"
          density="compact"
          v-model="minSalary"
          control-variant="hidden"
          label="Min Salary"
          type="number"
          :min="0"
        />
        <v-number-input
         color="primary"
         density="compact"
          v-model="maxSalary"
          control-variant="hidden"
          label="Max Salary"
          type="number"
        />
      </div>
    </div>
    <SkillsSelect :skills="skills" @update:skills="skills = $event"/>

    <ChipList v-if="skills.length > 0" :items="skills" @onClick="onClickChip" @onClean="clearSkills" label="Skills" maxHeight="150px"/>

  </form>
  
</template>
