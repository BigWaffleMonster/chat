<template>
  <div>
    chat
  </div>
</template>

<script lang="ts">
import store from '@/store'

export default {

  methods: {
    connectToWS() {
      console.log(store.state.nickname)
      }
  },

  mounted() {
    this.connectToWS()
  },
}
</script>

<style scoped>

</style>
