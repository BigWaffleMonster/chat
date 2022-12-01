<template>
  <div class="container flex items-center justify-center h-screen w-full">
    <div>
      <div class="bg-cyan-200 w-96 p-6 rounded shadow-sm">
        <div class="flex items-center justify-center mb-4">IMG</div>

        <label for="nickname">NickName</label>
        <my-input v-model="nickname" id="nickname" />
        <button type="submit" class="bg-cyan-300 w-full py-2 rounded hover:bg-cyan-400 transition-colors">
          Go
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      nickname: ''
    }
  }
}
</script>

<style scoped>

</style>
