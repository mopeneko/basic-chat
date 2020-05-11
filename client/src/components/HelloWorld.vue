<template>
  <v-container>
    <v-row justify="center" v-for="message in messages" :key="message.id">
      <v-col cols="12" md="8" lg="8" xl="8">
        <v-alert color="primary" outlined>{{ message.message }}</v-alert>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12" md="8" lg="8" xl="8">
        <v-text-field required :counter="140"></v-text-field>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" md="8" lg="8" xl="8">
        <v-btn color="primary" outlined>送信する</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "HelloWorld",

  data() {
    return {
      messages: []
    };
  },

  mounted() {
    const url = new URL(document.location.href);
    const origin = url.origin;
    const wsURL =
      origin.replace("http://", "ws://").replace("https://", "wss://") + "/ws";

    const socket = new WebSocket(wsURL);

    socket.addEventListener("open", () => {
      console.log("WebSocket is opened");
    });

    socket.addEventListener("message", event => {
      this.messages.push(event.data.message);
    });
  }
});
</script>
