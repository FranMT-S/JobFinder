<script setup lang="ts">
import { ref } from 'vue'
import { Seniority } from '../enums/seniority'



const props = withDefaults(defineProps<{
    seniority: Seniority
}>(), {
    seniority: Seniority.ANY
})


const seniorityOptions = ref([
  { value: Seniority.ANY, label: 'Any' },
  { value: Seniority.JUNIOR, label: 'Junior' },
  { value: Seniority.SEMISENIOR, label: 'Semi-Senior' },
  { value: Seniority.SENIOR, label: 'Senior' },
])
const selectedSeniority  = ref<Seniority>(props.seniority)
const emit = defineEmits(['update:seniority'])

const onChange = () => {
    emit('update:seniority', selectedSeniority.value);
}

</script>

<template>
    <v-select
      v-model="selectedSeniority"
      color="primary"
      :items="seniorityOptions"
      item-title="label"
      item-value="value"
      label="Seniority"
      density="compact"
      @update:modelValue="onChange"
    />
</template>


<style scoped>

</style>