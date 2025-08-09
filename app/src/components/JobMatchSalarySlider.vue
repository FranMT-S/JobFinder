<script setup lang="ts">
import { computed, ref } from 'vue';
import type { JobScrapeated } from '../types/Job';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';
import { useGlobalStore } from '../stores/globalStore';

  const props = defineProps<{
    jobScrapeated: JobScrapeated
    minSalary: number
    maxSalary: number
  }>()

  const sliderData = computed(() => {
    
    const KeyMaxJob = 'MaxJ'
    const KeyMinJob = 'MinJ'
    const KeyMaxExpected = 'MaxE'
    const KeyMinExpected = 'MinE' 
    const minSalary = props.minSalary
    const maxSalary = props.maxSalary
    const minSalaryJob = props.jobScrapeated.job.minimumSalary
    const maxSalaryJob = props.jobScrapeated.job.maximumSalary

    const ranges:Record<number, string> = {}
    const minMax:{min:number,max:number} = {min:0,max:0}
    const labels:Record<number, {icon:string,label:string}> = {}
    const values:{key:string,value:number}[] = [ {
          key: KeyMinJob,
          value: minSalaryJob
        },
        {
          key: KeyMaxJob,
          value: maxSalaryJob
        },
        {
          key: KeyMinExpected,
          value: minSalary
        },
        {
          key: KeyMaxExpected,
          value: maxSalary
        } ].filter(value => value.value > 0).sort((a, b) => a.value - b.value)

    if(minSalaryJob > 0 || maxSalaryJob > 0){
      let i = 0
      for(var value of values){
        if(value.key == KeyMinJob ){
          minMax.min = i
          minMax.max = i <= 0 ? i : minMax.max
        
        }

        if(value.key == KeyMaxJob ){
          minMax.max = i
          minMax.min = i <= 0 ? i : minMax.min
        }

        if(value.value <= 0 ) continue

        if(value.key == KeyMinExpected || value.key == KeyMaxExpected){
          ranges[i] = `${value.value /1000}K`
        }

        if(value.key == KeyMinJob || value.key == KeyMaxJob){
          ranges[i] = `${value.value /1000}K`
        }

        i++
      }
    }

    labels[minMax.min] = {icon:'',label:'Min'}
    labels[minMax.max] = {icon:'',label:'Max'}

    if(minSalary > 0 && minSalaryJob > minSalary){
      labels[minMax.min].icon = 'mdi-emoticon'
    }
    
    if(maxSalary > 0 && maxSalaryJob > maxSalary){
      labels[minMax.max].icon = 'mdi-emoticon'
    }


    return {ranges,minMax:[minMax.min,minMax.max],len:values.length,labels}
  })

const { isDarkMode } = storeToRefs(useGlobalStore())

</script>


<template>
   
  <v-range-slider
      readonly
      color="primary"
      v-model="sliderData.minMax"
      :min="0"
      :max="sliderData.len"
      hide-details
      show-ticks="always"
      thumb-label="always"
      step="1"
      :ticks="sliderData.ranges"
      :class="{ 'dark': isDarkMode }"
    >
      <template v-slot:thumb-label="{ modelValue }">
        <v-icon icon="mdi-emoticon"  color="primary" theme="dark" v-if="sliderData.labels[modelValue].icon"></v-icon>
        <span >{{ sliderData.labels[modelValue].label }}</span>
      </template>
    </v-range-slider>
</template>
  
<style scoped>
  :deep(.v-slider-thumb__label) {
    background-color: #FFFFFF;
  }

  :deep(.v-slider-thumb__label::before) {
    color: #FFFFFF;
  }

  .dark:deep(.v-slider-thumb__label) {
    background-color: #121212;  
  }

  .dark:deep(.v-slider-thumb__label::before) {
    color: #121212;
  }


  
</style>
  