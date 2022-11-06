import { onMounted, onUnmounted, ref } from 'vue';

export function useSwipe() {
	// register touch events

	onMounted(() => {
		document.addEventListener('touchstart', handleTouchStart, false);
		document.addEventListener('touchmove', handleTouchMove, false);
	});

	onUnmounted(() => {
		document.removeEventListener('touchstart', handleTouchStart, false);
		document.removeEventListener('touchmove', handleTouchMove, false);
	});

	let xDown = null;
	let yDown = null;

	let xDiff = ref(null);
	let yDiff = ref(null);

	const getTouches = (evt) => {
		return (
			evt.touches || // browser API
			evt.originalEvent.touches
		); // jQuery
	};

	const handleTouchStart = (evt) => {
		const firstTouch = getTouches(evt)[0];
		xDown = firstTouch.clientX;
		yDown = firstTouch.clientY;
	};

	const handleTouchMove = (evt) => {
		if (!xDown || !yDown) {
			return;
		}

		let xUp = evt.touches[0].clientX;
		let yUp = evt.touches[0].clientY;

		xDiff.value = xDown - xUp;
		yDiff.value = yDown - yUp;

		/* reset values */
		xDown = null;
		yDown = null;
	};

	return { xDiff, yDiff };
}
