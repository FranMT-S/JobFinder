<script setup lang="ts">
import { useGlobalStore } from '../stores/globalStore'
import { storeToRefs } from 'pinia'

const { isDarkMode, theme } = storeToRefs(useGlobalStore())
const { toggleDarkMode } = useGlobalStore()
</script>

<template>
  <v-btn icon @click="toggleDarkMode" size="small" variant="text">
    <v-icon :color="theme" size="large">
      {{ isDarkMode ? 'mdi-weather-night' : 'mdi-white-balance-sunny' }}
    </v-icon>
  </v-btn>
</template>

