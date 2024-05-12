<main
  class="min-h-screen bg-gray-800 flex flex-col items-center justify-center px-8 py-6 space-y-8"
>
  <h1 class="text-2xl font-extrabold text-white tracking-widest transition">
    Register
  </h1>

  <form class="w-full max-w-md space-y-6">
    <div class="w-full">
      <input
        class="w-full px-3 py-2 text-black rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500"
        type="email"
        placeholder="Email"
        required
      />
    </div>
    <div class="w-full">
      <input
        class="w-full px-3 py-2 text-black rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500"
        type="password"
        placeholder="Password"
        required
      />
    </div>
    <div class="w-full">
      <input
        class="w-full px-3 py-2 text-black rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500"
        type="password"
        placeholder="Confirm Password"
        required
      />
    </div>
    <button
      class="w-full text-lg font-bold text-white px-3 py-2 rounded-lg transform transition duration-500 ease-in-out hover:scale-105 hover:bg-green-500"
      type="submit"
    >
      Register
    </button>
  </form>
</main>
