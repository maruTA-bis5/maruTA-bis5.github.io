(function() {
    if (typeof ogpService !== 'undefined') {
        convertAnchorToCard();
    }
})();

function convertAnchorToCard() {
    const anchors = document.querySelectorAll("a.createCard");
    anchors.forEach(convertAnchor);
}

async function convertAnchor(anchor) {
    await createCard(anchor);
    anchor.classList.toggle("createCard");
}

async function createCard(anchor) {
    const href = anchor.attributes.getNamedItem("href").value;
    fetch(ogpService + "?url=" + href)
        .then(response => response.json())
        .then(ogp => {
            if (ogp.title) {
                const card = createCardFromJson(ogp);
                insertCard(anchor, card);
            }
        });
}

function createCardFromJson(ogp) {
    const title = ogp.title;
    const desc = ogp.description;
    const url = ogp.url;
    const siteName = ogp.site_name;
    const image = ogp.image_url;

    const containerNode = createDiv('link-card');
    const innerContainer = createDiv('link-card-container');
    containerNode.appendChild(innerContainer);

    const textContainer = createDiv('link-card-text-container');
    innerContainer.appendChild(textContainer);

    if (siteName) {
        const siteNameEl = createDiv('link-card-site-name');
        siteNameEl.appendChild(document.createTextNode(`${siteName}`));
        textContainer.appendChild(siteNameEl);
    }

    const titleEl = createDiv('link-card-title');
    const titleAnchor = document.createElement('a');
    titleAnchor.setAttribute('href', `${url}`);
    titleAnchor.appendChild(document.createTextNode(`${title}`));
    titleEl.appendChild(titleAnchor);
    textContainer.appendChild(titleEl);

    if (desc) {
        const descEl = createDiv('link-card-description');
        descEl.appendChild(document.createTextNode(`${desc}`));
        textContainer.appendChild(descEl);
    }

    if (image) {
        const imageEl = document.createElement('img');
        imageEl.setAttribute('alt', 'image');
        imageEl.setAttribute('src', `${image}`);
        innerContainer.appendChild(imageEl);
    }

    // TODO escape to avoid xss 
    return containerNode;
}

function createDiv(className) {
    const el = document.createElement('div');
    el.classList.add(className);
    return el;
}

function insertCard(anchor, card) {
    anchor.parentNode.insertBefore(card, anchor.nextSibling);
}
