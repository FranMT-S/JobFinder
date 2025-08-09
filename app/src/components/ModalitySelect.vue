<script setup lang="ts">
import { ref } from 'vue'
import { Modality } from '../enums/modality'
import { storeToRefs } from 'pinia'
import { useUserStore } from '../stores/userStore'

const  {modality} = storeToRefs(useUserStore())

const modalityOptions = ref([
  { value: Modality.REMOTE, label: 'Remote' },
  { value: Modality.HYBRID, label: 'Hybrid' },
  { value: Modality.PRESENTIAL, label: 'Presential' },
  { value: Modality.ANY, label: 'Any' },
])

</script>

<template>
    <v-select
      v-model="modality"
      density="compact"
      color="primary"
      :items="modalityOptions"
      item-title="label"
      item-value="value"
      label="Modality"
    />
</template>


<style scoped>

</style>