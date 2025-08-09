
<script setup lang="ts">
import { ref } from 'vue';
import { useGlobalStore } from '../stores/globalStore';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';
import { computed } from '@vue/reactivity';


const {skillsList} = storeToRefs(useGlobalStore())
const {addSkill, clearSkills} = useUserStore()
const selected = ref<string | null>(null)
const refInput = ref()

const {skills:currentSkills} = storeToRefs(useUserStore())
const skillsMap = computed(() => new Set<string>(currentSkills.value.map(s => s.toLowerCase())))

const capitalizeSkills = computed(() => {
  const capitalizedSkills: string[] = []
  for(let i = 0; i < skillsList.value.length; i++){
    const skill = skillsList.value[i]
    if(skillsMap.value.has(skill.toLowerCase()))
      continue

    const capitalizeSkills = skill[0].toUpperCase() + skill.slice(1)
    capitalizedSkills.push(capitalizeSkills)
  }
  return capitalizedSkills
})

const onKeypress = (e: KeyboardEvent) => {
  const input = refInput.value.$el.querySelector('input')
  if(!input)
    return

  const valueInput = input.value.trim()
  if (valueInput) {
    addSkill(valueInput)
    clearInput()
    setTimeout(() => {
      focusInput()
    }, 0)
  }
}

const onSelectSkill = (e:string) =>{
  if(!e || e.trim() === '')
    return

  addSkill(e)
  clearInput()
  setTimeout(() => {
      focusInput()
  }, 0)
}

const clearInput = () => {
  const input = refInput.value.$el.querySelector('input')
  if(!input)
    return

  selected.value = null
  input.value = ''
  input.blur()
}

const focusInput = () => {
  const input = refInput.value.$el.querySelector('input')
  if(!input)
    return

  refInput.value.$el.focus()
  input.focus()
}

</script>

<template>
  <v-autocomplete
    ref="refInput"
    :model-value="selected"
    :items="capitalizeSkills"
    label="Skills"
    variant="outlined"
    density="compact"
    color="primary"
    bg-color="surfaceT"
    :clear-on-select="true"
    :hide-selected="true"
    @keydown.enter="onKeypress"
    @click:clear="clearSkills"
    @update:modelValue="onSelectSkill"
  />
</template>


