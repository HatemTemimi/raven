<script setup>
</script>

<template>
  <div class="container" style="width: 100vw; height: 100vh">
    <div  style="height: 100%; display: flex; flex-direction: column; justify-content: space-between">
      <div style="text-align: center; height: 5%" class="bordered">navbar</div>
      <div style="height: 80%" class="bordered">
        <div style="padding: 4px; height: 100%" class="bordered">
          <div>
            the proxy table
          </div>
          <div>
            table data
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.bordered {
  border-width: 2px;
  border-color: black;
  border-style: solid;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
