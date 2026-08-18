const share = async (data: ShareData) => {
  await navigator.share(data);
};

const buttons = document.querySelectorAll('button');

buttons.forEach((el) => {
  const title = el.getAttribute('data-share-title') || 'Blog Post';
  const url = el.getAttribute('data-share-url') || document.location.toString();

  el.addEventListener('click', () => share({
    title,
    text: title,
    url,
  }));
})
