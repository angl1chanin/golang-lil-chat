<template>
  <div class="chat app-window">
    <div class="chat__wrapper wrapper">
      <p class="chat__title title">Chat</p>
      <Preloader v-if="isMessagesLoading" />
      <ul v-else-if="messages.length > 0" class="chat__messages list">
        <li class="chat__message" v-for="message in messages" :key="message.id">
          <div class="chat__message-wrapper">
            <Message :message="message" />
          </div>
        </li>
      </ul>
      <p v-else class="chat__empty empty">No messages yet</p>
      <div class="chat__send" v-if="!isMessagesLoading">
        <form
            class="chat__send-form app-form app-window"
            action=""
            method="post"
            @submit.prevent="sendForm"
        >
          <input
              class="app-form__input"
              type="text"
              placeholder="Username"
              v-model="author"
          />
          <div style="display: flex">
            <input
                class="app-form__input"
                type="text"
                placeholder="Message"
                v-model="message"
            />
            <button class="app-form__button button" type="submit">
              <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="14"
                  height="14"
                  viewBox="0 0 24 24"
                  fill="none"
              >
                <path
                    d="M10.3009 13.6949L20.102 3.89742M10.5795 14.1355L12.8019 18.5804C13.339 19.6545 13.6075 20.1916 13.9458 20.3356C14.2394 20.4606 14.575 20.4379 14.8492 20.2747C15.1651 20.0866 15.3591 19.5183 15.7472 18.3818L19.9463 6.08434C20.2845 5.09409 20.4535 4.59896 20.3378 4.27142C20.2371 3.98648 20.013 3.76234 19.7281 3.66167C19.4005 3.54595 18.9054 3.71502 17.9151 4.05315L5.61763 8.2523C4.48114 8.64037 3.91289 8.83441 3.72478 9.15032C3.56153 9.42447 3.53891 9.76007 3.66389 10.0536C3.80791 10.3919 4.34498 10.6605 5.41912 11.1975L9.86397 13.42C10.041 13.5085 10.1295 13.5527 10.2061 13.6118C10.2742 13.6643 10.3352 13.7253 10.3876 13.7933C10.4468 13.87 10.491 13.9585 10.5795 14.1355Z"
                    stroke="#000000"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                />
              </svg>
            </button>
          </div>
        </form>
        <button class="chat__send-button button" @click="toggleForm">
          New message
        </button>
      </div>
    </div>
  </div>
</template>

<script>
// import the styling for the toast
import "mosha-vue-toastify/dist/style.css";

// components
import Message from "@/components/Message";
import Preloader from "@/components/Preloader";

// libraries
import notices from "@/mixins/notices";

export default {
  name: "Chat",
  components: {
    Message,
    Preloader,
  },
  mixins: [notices],
  data() {
    return {
      messages: [],
      isFormOpened: false,
      isMessagesLoading: false,
      author: "",
      message: "",
      api: "http://localhost:8081/v1/chat/messages",
    };
  },
  methods: {
    loadMessages() {
      this.isMessagesLoading = true;
      fetch(this.api)
          .then((response) => {
            if (response.ok) {
              return response.json();
            } else {
              this.danger("Error while loading chat");
            }
          })
          .then((json) => {
            this.messages = json;
            this.isMessagesLoading = false;
          })
          .catch(() => {
            console.log("Chat not available");
          });
    },
    toggleForm() {
      let form = document.querySelector(".chat__send-form");
      if (this.isFormOpened) {
        form.style.top = "0";
        form.style.opacity = "0";
        setTimeout(() => {
          form.style.visibility = "hidden";
        }, 200);
      } else {
        form.style.top = "-100%";
        form.style.opacity = "1";
        form.style.visibility = "visible";
      }
      this.isFormOpened = !this.isFormOpened;
    },
    async sendForm() {
      if (this.author !== "" && this.message !== "") {
        const data = {
          author: this.author,
          text: this.message,
        };

        const config = {
          method: "POST",
          body: JSON.stringify(data),
        };

        await fetch(this.api, config)
            .then((response) => {
              return response.json();
            })
            .then((json) => {
              if (json.status === "success") {
                this.success("Message successfully sent");
                this.loadMessages();
              } else {
                this.danger(json.error);
              }
            })
            .catch(() => {
              this.danger("Chat server isn't available");
            });
        this.author = "";
        this.message = "";
      } else {
        this.warning("All fields are required");
      }
    },
  },
  mounted() {
    this.loadMessages();
  },
};
</script>

<style lang="scss" scoped>
@mixin scrollbar {
  &::-webkit-scrollbar {
    width: 1px;
  }

  &::-webkit-scrollbar-thumb {
    background-color: #999;
  }
}

.app-window {
  background-color: #fff;
  padding: 30px;
  border-radius: 32px;
  box-shadow: 0 2px 8px 0 rgba(34, 60, 80, 0.2);
}

.app-form {
  box-shadow: 0 4px 8px 0 rgba(34, 60, 80, 0.2);
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding: 10px;
  border-radius: 16px;

  &__input {
    border: none;
    outline: none;
    font-size: 14px;
    font-family: inherit;
  }

  &__button {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 35px;
    height: 35px;
  }
}

.button {
  background: rgba(118, 118, 128, 0.12);
  color: var(--text-color);
  border: none;
  font-family: "Poppins", sans-serif;
  font-size: 16px;
  font-weight: 400;
  text-align: center;
  border-radius: 12px;
  line-height: 22px;
  cursor: pointer;
  padding: 10px 0;
}

.title {
  color: var(--text-color);
  font-size: 18px;
  font-weight: 500;
  text-transform: capitalize;
}

.chat {
  aspect-ratio: 2/1;

  &__title {
    margin-bottom: 35px;
  }

  &__empty {
    text-align: center;
    margin-bottom: 35px;
  }

  &__send-button {
    width: 100%;
  }

  &__messages {
    max-height: 140px;
    display: flex;
    flex-direction: column;
    padding: 0;
    margin: 0 0 35px 0;
    overflow-y: auto;

    @include scrollbar;
  }

  &__message {
    list-style-type: none;
    cursor: default;
    &:not(:last-child) {
      margin-bottom: 30px;
    }
  }

  &__message-wrapper {
    display: flex;
    align-items: flex-start;
  }

  &__send {
    position: relative;
  }

  &__send-form {
    position: absolute;
    top: 0;
    opacity: 0;
    visibility: hidden;
    left: 50%;
    transform: translate(-50%, -50%);
    transition: top 0.2s linear, opacity 0.2s linear;
  }
}
</style>