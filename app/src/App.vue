<script setup lang="ts">
import { onBeforeMount, onMounted, ref } from 'vue';
import JobSeekerForm from './components/JobSeekerForm.vue';
import ToastBanner from './components/ToastBanner.vue';
import { useToast } from './composable/useToast'
import { storeToRefs } from 'pinia';
import { useJobsStore } from './stores/JobsStore';
import JobMatch from './components/JobMatch.vue';
import { useGlobalStore } from './stores/globalStore'
import { useUserStore } from './stores/userStore';
import { fetchScrapper } from './services/fetchScrapper';
import ToggleDarkModeButton from './components/ToggleDarkModeButton.vue';

const {message, show, trigger } = useToast()
const {jobScrapeated, isLoading, firstLoad} = storeToRefs(useJobsStore())
const {initAllData} = useGlobalStore()
const { theme,isSalaryRange,hostList,maxPage } = storeToRefs(useGlobalStore())
const { location, level, skills, modality, minSalary, maxSalary, category,hostSelected } = storeToRefs(useUserStore())

let currentController: AbortController | null = null;

const handleSubmit = async () => {
  firstLoad.value = true
  isLoading.value = true
  if(currentController){
    currentController.abort()
    currentController = null
  }

  const minSalaryValue = isSalaryRange.value ? minSalary.value : 0
  const maxSalaryValue = isSalaryRange.value ? maxSalary.value : 0

  currentController = new AbortController()
  const signal = currentController.signal
  

  const [err,result] = await fetchScrapper(1,maxPage.value,{
    location: location.value,
    level: level.value,
    skills: skills.value,
    modalities: modality.value,
    minimumSalaryExpectation: minSalaryValue,
    maximumSalaryExpectation: maxSalaryValue,
    position: '',
    category: category.value,
    hostSelected: hostSelected.value
  },signal)

  if((err instanceof DOMException && err.name === "AbortError")){
    return
  }

  if(err){
    firstLoad.value = false
    isLoading.value = false
    trigger(err.message)
    return
  }

  jobScrapeated.value = result
  isLoading.value = false
}

onBeforeMount(async () => {
  await initAllData()
})

const toggleHost = (hostId: number) => {
  if(hostSelected.value.includes(hostId)){
    hostSelected.value = hostSelected.value.filter(id => id !== hostId)
    return
  }

  hostSelected.value.push(hostId)
}

</script>

<template>
 <v-app :theme="theme" >
  <v-layout class="h-screen">
    <v-navigation-drawer
      app
      permanent
      color="surface"
      class="overflow-y-auto"
      style="height: 100vh;"
    >
      <v-list class="text-end px-2">
        <v-row class="ps-2">
          <v-col cols="8" md="8" class="flex items-end   gap-2">
            <v-icon icon="mdi-briefcase-search-outline" color="primary" />
            <v-label>Max Page</v-label>
          </v-col>
          <v-col cols="4" md="4">
            <ToggleDarkModeButton />
          </v-col>
        </v-row>
       
       <v-list-item class="!p-0 mt-2">
          <v-slider
            v-model="maxPage"
            :max=100
            :min=1
            step=1
            class="align-center ms-5"
            hide-details
          >
            <template v-slot:append>
              <v-text-field
                :max="100"
                :min="1"
                step=1
                label="Max"
                control-variant="hidden"
                class="no-spinner"
                v-model="maxPage"
                density="compact"
                style="width: 70px"
                type="number"
                hide-details
              ></v-text-field>
            </template>
          </v-slider>
        </v-list-item>  
      </v-list>
      


      <v-divider></v-divider>

      <v-list>
        <v-list-item>
         <v-label class="mb-2">Host To Scrap</v-label>
         <div class="flex flex-wrap gap-2">
          <v-chip
            size="small"
            v-for="host in hostList"
            :key="host.id"
            :color="hostSelected.includes(host.id) ? 'active' : 'inactive'"
            @click="toggleHost(host.id)"
          >
            {{ host.name }}
          </v-chip>
         </div>
        </v-list-item>
        <v-list-item>
          <p class="text-[1.0rem] text-primary ">Enter your data</p>
        </v-list-item>
        <v-list-item>
          <JobSeekerForm />
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn color="primary" block @click="handleSubmit">
            Search
          </v-btn>
        </div>
      </template>

    </v-navigation-drawer>

    <v-main class="overflow-y-auto " style="height: 100vh;">
      <v-container >
        <ToastBanner :message="message" :show="show"/>
        <v-container v-if="isLoading" >
          <v-row v-for="key in [1,2,3,4]" :key="'skeleton' + key">
            <v-col cols="12" md="12" elevation="10">
              <v-skeleton-loader
                type="chip,table-heading,text@2,text@3"
              />
            </v-col>
          </v-row>
        </v-container>
        
        
        <div v-if="jobScrapeated.length == 0 && !isLoading && firstLoad" class="flex justify-center items-center h-full">
          <v-empty-state
            class="flex justify-center items-center h-full"
            icon="mdi-magnify"
            text="Try adjusting your search terms or filters. Sometimes less specific terms or broader queries can help you find what you're looking for."
            title="We couldn't find a match."
          />
        </div>
  
        <JobMatch v-if="jobScrapeated.length > 0 && !isLoading" v-for="job in jobScrapeated" :key="job.job.url" :jobScrapeated="job" />
      </v-container>
    </v-main>
  </v-layout>
</v-app>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
