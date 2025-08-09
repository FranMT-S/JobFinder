
<script setup lang="ts">
  import type { MatchAnalizer } from '../types/Job';
  import { onBeforeUnmount, ref } from 'vue';
  
  defineProps<{
    matchAnalizer: MatchAnalizer
  }>()

  const isShow = ref(false)

  const refCard = ref()

  const clickOutside = (event: MouseEvent) => {
    if(!isShow.value || !refCard.value){
      return
    }

    const div = refCard.value.$el

    if (!div.contains(event.target as Node)) {
      isShow.value = false
    } 
  }

  const toggleShow = (event: MouseEvent) => {
    event.stopPropagation()
    isShow.value = !isShow.value
    if(isShow.value){
      document.addEventListener('click', clickOutside)
    }else{
      document.removeEventListener('click', clickOutside)
    }
  }

  onBeforeUnmount(() => {
    document.removeEventListener('click', clickOutside)
  })
</script>

<template>
  <div class="relative">
    <v-chip color="purple-lighten-1" @click="(e:any) => toggleShow(e)" :title="isShow ? 'Hide' : 'Show'">
      <span class="me-1">Match</span>  <span class="font-weight-bold">{{ matchAnalizer.totalPorcent }}%</span>
    </v-chip>
    <v-card color="surface" ref="refCard" class="position-absolute !p-4 top-full left-0  !z-[999999999999]" v-if="isShow">
      <div class="text-capitalize flex gap-10 w-max">
        <ul>
          <li class="mb-1">
            <v-icon size="small" color="primary" icon="mdi-map-marker-radius" />
            modalities: {{ matchAnalizer.porcentModalities }}%
          </li>
          <li class="mb-1">
            <v-icon size="small" color="primary" icon="mdi-account" />
            levels: {{ matchAnalizer.porcentLevels }}%
          </li>
          <li class="mb-1">
            <v-icon size="small" color="primary" icon="mdi-briefcase" />
            position: {{ matchAnalizer.porcentPosition }}%
          </li>
        </ul>
        <ul>
          <li class="mb-1">
            <v-icon size="small" color="primary" icon="mdi-currency-usd" />
            salary: {{ matchAnalizer.porcentSalary }}%
          </li>
          <li class="mb-1">
            <v-icon size="small" color="primary" icon="mdi-cog" />
            skills: {{ matchAnalizer.porcentSkills }}%
          </li>
        </ul>
      </div>
      </v-card>
  </div>
</template>