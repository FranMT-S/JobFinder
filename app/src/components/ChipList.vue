
<script setup lang="ts">
import { ref } from 'vue'

const props = withDefaults(defineProps<{
  items: Array<string>
  label?: string
  maxHeight?: string
}>(), {
  items: () => [],
  label: 'Yours Skills'
})

const emit = defineEmits(['onClick','onClean'])


const removeChip = (index: number) => {
  emit('onClick', index)
}

const clearAll = () => {
  emit('onClean')
}

</script>

<template>
  <div class="mb-2 " >
    
    <div class="flex items-center justify-between">
      <v-label class="mb-1">{{ props.label }}</v-label>
      <v-icon v-if="props.items.length > 0" icon="mdi-trash-can-outline cursor-pointer" @click="clearAll"></v-icon>
    </div>
    <div class="flex flex-wrap gap-2 overflow-y-auto" :style="{ 'max-height': maxHeight }">
      <v-chip
        color="primary"
        class="font-weight-medium capitalize"
        v-for="(item, index) in props.items"
          :key="index"
          @click="removeChip(index)"
        >
          {{ item }}
        </v-chip>
    </div>
  </div>
</template>
