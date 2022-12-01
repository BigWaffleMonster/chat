<template>
  <input :value="modelValue" @input="updateInput" type="text"
    class="w-full py-2 bg-emerald-300 px-1 outline-none mb-4 rounded" />
</template>

<script lang="ts">
export default {
  name: 'my-input',
  props: {
    modelValue: String
  },
  methods: {
    updateInput(event: any) {
      this.$emit('update:modelValue', event.target.value)
    }
  }
}
</script>

<style scoped>

</style>
