<template>
  <div class="chat__message-icon">
    <div class="chat__message-icon-wrapper">
      <img :src="uiAvatar" alt="">
    </div>
  </div>
  <div class="chat__message-body">
    <p class="chat__message-username">{{ message.username }}</p>
    <p class="chat__message-message">{{ message.text }}</p>
  </div>
  <div class="chat__message-info">
    <p class="chat__message-date">{{ dateFromTimestamp }}</p>
  </div>
</template>

<script>
export default {
  name: "Message",
  props: {
    message: {
      type: Object,
      required: true,
    }
  },
  computed: {
    dateFromTimestamp() {
      return this.$props.message.timestamp.slice(0,10).replaceAll("-","/")
    },
    uiAvatar() {
      const bg = [
          "0288D1",
          "0097A7",
          "00796B",
          "388E3C",
          "689F38",
          "AFB42B",
          "FBC02D",
          "FFA000",
          "F57C00",
          "E64A19"
      ]
      return `https://ui-avatars.com/api/?name=${this.$props.message.username}&background=${bg[Math.floor(Math.random()*bg.length)]}&color=fff`
    }
  }
}
</script>

<style lang="scss" scoped>
.chat {
  &__message-icon {
     margin-right: 12px;
   }

  &__message-icon-wrapper img {
     width: 32px;
     height: 32px;
     border-radius: 100%;
   }

  &__message-username {
     color: var(--text-color);
     font-weight: 500;
     margin-bottom: 12px;
   }

  &__message-message {
     word-break: break-all;
   }

  &__message-info {
     margin-left: auto;
     margin-right: 2px;
   }

  &__message-date {
     color: #999999;
     font-size: 12px;
   }
}
</style>