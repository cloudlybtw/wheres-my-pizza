# 🍴 Restaurant Order Processing System 🍕🍔

Welcome to the Restaurant Order Processing System! 🎉
This is a fun, simple, and *super cute* project that simulates the operations of a bustling restaurant kitchen! 🧑‍🍳👩‍🍳

### ✨ What is it?

This system is made up of three key services that work together to make sure your food gets cooked, tracked, and served super smoothly:

1. Kitchen Worker 🧑‍🍳
   The heart of the kitchen! These workers process incoming orders, cook them (with the help of some simulated kitchen magic ✨), and update the status of each order.

2. Tracking Service 📡
   Curious about where your order is? This service lets you track the status of your orders and see a history of all the steps your food has gone through.

3. Notification Service 📲
   Get real-time notifications on order status updates! Whether it's your chef moving things into the "cooking" phase or your order being ready for pickup, this service makes sure you're always in the loop! 💬

---

### 🍕 How does it work?

* Kitchen Worker:

  * Picks up orders from RabbitMQ 🐇, cooks 'em (or simulates it! ⏳), and updates the order status from *received* → *cooking* → *ready*.
  * The workers are super diligent! They also let you know when their shift is over by updating their status to *offline*. ✌️

* Tracking Service:

  * You can ask it where your food is with simple HTTP requests like:

    * "Hey, where’s my order?" (GET /orders/{order\_number}/status)
    * "Tell me the full history of my order!" (GET /orders/{order\_number}/history)
    * "What’s the status of the chefs?" (GET /workers/status)

* Notification Service:

  * Whenever an order's status changes (like from *received* to *cooking*), everyone is notified instantly! 💌
  * No matter how many subscribers you have, all of them get the update. It's like a kitchen gossip train that just keeps on chugging! 🚂💬

---

### 🏃‍♀️ How to Run It?

#### 1. Start a Kitchen Worker

To start a general worker:

./restaurant-system --mode=kitchen-worker --worker-name="chef_anna" --prefetch=1
Start a specialized worker (e.g., only handling dine-in orders):

./restaurant-system --mode=kitchen-worker --worker-name="chef_mario" --order-types="dine_in" &
#### 2. Start the Tracking Service

Track all orders and workers in real time with this API service:

./restaurant-system --mode=tracking-service --port=3002
#### 3. Start the Notification Service

Get notifications like your favorite kitchen-themed show:

./restaurant-system --mode=notification-subscriber
---

### 🎉 Features You’ll Love

* Super Cute Worker Names: Choose fun names for your workers (Chef Mario, Chef Luigi, anyone? 🍄)
* Track Your Order Like a Pro: See exactly where your food is in the process. Is it cooking? Or just waiting to be ready for you? 🕒
* Notifications for Days: Get updates in real time! It's like you're sitting right in the kitchen with the chefs 👩‍🍳💬

---

### 💬 Logs & Debugging

If something goes wrong (but let's hope it doesn't! 😅), you'll get detailed logs, including:

* INFO: Worker registration, service startup, and status updates.
* DEBUG: What exactly is going on in the background.
* ERROR: If something breaks, we’ll tell you exactly where.

---

### 🛠 Tech Used

* RabbitMQ 🐰: For message queuing (because kitchens are all about communication!)
* PostgreSQL 🍲: For keeping track of orders and workers.
* HTTP API: To keep everything connected to the outside world.
* Node.js 🚀: The backbone of our project.

---

### 🧑‍🍳 Special Thanks

To all the hardworking chefs (aka developers) who make sure this system runs smoothly and keeps the kitchen buzzing! 🍽️

---

### 💖 Contribute

If you have ideas to make the restaurant even better, feel free to fork this repo and send a pull request! We love improvements, just like we love a perfectly cooked meal. 🍳

---

### 🍴 Bon Appétit!

Enjoy your food, your order, and your notifications! We hope this system serves you well, whether you're working in the kitchen or just enjoying your meal. 😄