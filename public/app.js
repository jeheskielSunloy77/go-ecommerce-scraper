const search = () => {
	const loader = `<div id="loader" class="absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
    <div role="status" class="flex flex-col items-center justify-center">
        <svg class="inline mr-2 w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
            <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
        </svg>
        <span class="text-gray-700 dark:text-gray-200">Crawling the web for product...</span>
    </div>
</div>`
	document.querySelector('#search-status').innerHTML = loader
	const searchParam = document.querySelector('input').value
	const fetchData = async () => {
		const response = await fetch(`/api/json?q=${searchParam}`)
		const data = await response.json()
		return data
	}
	fetchData().then((data) => {
		document
			.querySelector('#search-status')
			.innerHTML = `<span class="text-gray-700 dark:text-gray-200 mb-4">${data.length} total item found!</span>`
		const markup = data.map((item) => {
			return `
<div class="w-full bg-gray-50 rounded-lg shadow-md hover:shadow-lg hover:dark:shadow-2xl dark:bg-gray-800 dark:border-gray-700 flex flex-col content-between">
    <a href=${item.productUrl} >
        <img class="p-8 rounded-t-lg object-contain mx-auto" src=${item.imageUrl} alt="product image" />
    </a>
    <div class="px-5 pb-5">
        <a href=${item.productUrl} >
            <h5 class="font-semibold tracking-tight text-gray-900 dark:text-white">${item.name}</h5>
        </a>
        <span class="text-gray-500 dark:text-gray-400">${item.vendor}</span>
        <div class="flex items-center justify-between">
            <span class="text-3xl font-bold text-gray-900 dark:text-white">${item.price}</span>
            <a href=${item.productUrl}  class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">View Product</a>
        </div>
    </div>
</div>`

		})
		document
			.querySelector('#search-result')
			.innerHTML = markup.join('')
		document.querySelector('#loader').remove()
	})
}

var themeToggleDarkIcon = document.getElementById('theme-toggle-dark-icon')
var themeToggleLightIcon = document.getElementById('theme-toggle-light-icon')

if (
	localStorage.getItem('color-theme') === 'dark' ||
	(!('color-theme' in localStorage) &&
		window.matchMedia('(prefers-color-scheme: dark)').matches)
) {
	themeToggleLightIcon.classList.remove('hidden')
} else {
	themeToggleDarkIcon.classList.remove('hidden')
}

var themeToggleBtn = document.getElementById('theme-toggle')

themeToggleBtn.addEventListener('click', function () {
	themeToggleDarkIcon.classList.toggle('hidden')
	themeToggleLightIcon.classList.toggle('hidden')

	if (localStorage.getItem('color-theme')) {
		if (localStorage.getItem('color-theme') === 'light') {
			document.documentElement.classList.add('dark')
			localStorage.setItem('color-theme', 'dark')
		} else {
			document.documentElement.classList.remove('dark')
			localStorage.setItem('color-theme', 'light')
		}
	} else {
		if (document.documentElement.classList.contains('dark')) {
			document.documentElement.classList.remove('dark')
			localStorage.setItem('color-theme', 'light')
		} else {
			document.documentElement.classList.add('dark')
			localStorage.setItem('color-theme', 'dark')
		}
	}
})
