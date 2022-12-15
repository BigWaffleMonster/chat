<template>
  <div>
    chat
  </div>
</template>

<script lang="ts">
import store from '@/store'

export default {
  data() {
    return {
      conneted: false,
      socket: null,
      messages: []
    }
  },

  mounted() {
    this.socket = new WebSocket('ws://localhost:8080')

    this.socket.onopen = () => {
      this.conneted = true

      const message = {
        event: 'connection',
        username: store.state.nickname,
        id: Date.now(),
        body: `${store.state.nickname} has connected`
      }

      this.socket.send(JSON.stringify(message))
    }

    this.socket.onmessage = (event) => {
      const message = JSON.parse(event.data)

      console.log(message)
    }

    this.socket.onerror = () => {
      console.log('Error')
    }

    this.socket.onclose = () => {
      console.log('socket has been closed')
    }
  },
}
</script>

<style scoped>

</style>
