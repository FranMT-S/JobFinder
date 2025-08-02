<script setup lang="ts">
  const props = defineProps({ message: String, show: Boolean })

</script>

<template>
    <transition name="fade">
      <div v-if="props.show" class="fixed top-4 right-4 z-50 bg-red-600 text-white px-4 py-2 rounded shadow">
        {{ props.message }}
      </div>
    </transition>
</template>

  
  
<style scoped>
.fade-enter-active,
.fade-leave-active {
transition: opacity 0.5s;
}
.fade-enter-from,
.fade-leave-to {
opacity: 0;
}
</style>