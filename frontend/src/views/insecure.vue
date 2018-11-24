<template>
  <section>
    <div id="loggedIn" class="container is-fluid">
      <section>
        <video class="center" width="50%" height="480" autoplay loop style="z-index:1000000; margin-left:20%; position: fixed; border-radius: 5px; margin-top: 1%">
          <source src="../assets/no.mp4" type="video/mp4" />
        </video>
      </section>
    </div>
    <vue-particles>
      color="#FFFFFF"
      :particleOpacity="0.7"
      :particlesNumber="200"
      shapeType="circle"
      :particleSize="4"
      linesColor="#dedede"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="false"
      :clickEffect="false"
    &gt;
    </vue-particles>
  </section>
</template>

<script>
  export default {
    name: "Insecure",
    data() {
      return {};
    }
  };
</script>

<style scoped>
  #insecure {
    background-color: #ffffff;
    border: 1px solid #cccccc;
    padding: 20px;
    margin-top: 1%;
  }
  .center {
    margin-left: auto;
    margin-right: auto;
    position: relative;
    display: block;
  }
</style>