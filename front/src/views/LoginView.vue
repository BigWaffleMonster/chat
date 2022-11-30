<template>
  <div class="container flex items-center justify-center h-screen">
    <div>
      <div class="bg-cyan-200 w-96 p-6 rounded shadow-sm">
        <div class="flex items-center justify-center mb-4">IMG</div>

        <label for="nickname">NickName</label>
        <input 
          type="text"
          id="nickname"
          class="w-full py-2 bg-emerald-300 px-1 outline-none mb-4 rounded"
        />

        <button 
          type="submit" 
          class="bg-cyan-300 w-full py-2 rounded hover:bg-cyan-400 transition-colors"
        >
          Go
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
</script>

<style scoped>

</style>
