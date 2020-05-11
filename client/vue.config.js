module.exports = {
  transpileDependencies: ["vuetify"],

  devServer: {
    proxy: {
      "/ws": {
        target: "ws://localhost:4000",
        changeOrigin: true
      }
    }
  }
};
