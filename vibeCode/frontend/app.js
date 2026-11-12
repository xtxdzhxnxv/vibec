const navButtons = document.querySelectorAll('button[data-page]');
const pages = document.querySelectorAll('.page');

const isLoggedIn = false;


navButtons.forEach(button => {
  button.addEventListener('click', () => {
    const targetPage = document.getElementById(`${button.dataset.page}-page`);

    if (!targetPage) {
      console.error(`Страница "${button.dataset.page}" не найдена`);
      return;
    }

    if (targetPage === document.querySelector('.page.account-page')) {
        const accountPage = document.querySelector('.page.account-page');
        const catalogPage = document.querySelector('.page.is-visible');
        if (isLoggedIn === false) {
            accountPage.hidden = true; 
            const authDialog = document.querySelector('.auth-dialog');
            authDialog.style.display = 'block';
            authDialog.classList.add('auth-dialog-active');
            catalogPage.hidden = true;
          
        
        document.querySelector('#close-auth-dialog').addEventListener('click', () => {
            const authDialog = document.querySelector('.auth-dialog');
            authDialog.style.display = 'none';
            authDialog.classList.remove('auth-dialog-active');
            catalogPage.hidden = false;
          });
          return;
        }
        
    }
    pages.forEach(page => {
      const isTarget = page === targetPage;
      page.hidden = !isTarget;
      page.classList.toggle('is-visible', isTarget);
    });

    document.querySelectorAll('.nav-item[data-page]').forEach(item => {
        const isActive = item === button;
        item.classList.toggle('is-active', isActive);
        openListingDialogButton.classList.remove('is-active');
        listingDialog.style.display = 'none';
        document.querySelector('.auth-dialog').style.display = 'none';

      if (isActive) {
        item.setAttribute('aria-current', 'page');
      } else {
        item.removeAttribute('aria-current');
      }
    });
  });
})

const listingDialog = document.getElementById('listing-dialog');
const openListingDialogButton = document.getElementById('bottom-add-button');
const closeListingDialogButton = document.getElementById('close-dialog');
const catalogPage = document.getElementById('catalog-page');
const accountPage = document.getElementById('account-page');

//кнопка создания объявления
if (isLoggedIn === true) {
openListingDialogButton.addEventListener('click', () => {
    listingDialog.style.display = 'block';
    catalogPage.hidden = true;
    accountPage.hidden = true;
    openListingDialogButton.classList.add('is-active');
    document.querySelector('.nav-item[data-page="catalog"]').classList.remove('is-active');
    document.querySelector('.nav-item[data-page="account"]').classList.remove('is-active');   

});

closeListingDialogButton.addEventListener('click', () => {
    listingDialog.style.display = 'none';
    catalogPage.hidden = false;
    accountPage.hidden = true;
    openListingDialogButton.classList.remove('is-active');
    document.querySelector('.nav-item[data-page="catalog"]').classList.add('is-active');
    document.querySelector('.nav-item[data-page="account"]').classList.remove('is-active');
});

};


//хэдер кнопка входа
document.querySelector('#header-auth-button').addEventListener('click', () => {
    const accountPage = document.querySelector('.page.account-page');
        const catalogPage = document.querySelector('.page.is-visible');
            accountPage.hidden = true; 
            const authDialog = document.querySelector('.auth-dialog');
            authDialog.style.display = 'block';
            catalogPage.hidden = true;
            listingDialog.style.display = 'none';


            document.querySelector('#close-auth-dialog').addEventListener('click', () => {
                const authDialog = document.querySelector('.auth-dialog');
                authDialog.style.display = 'none';
                listingDialog.style.display = 'none';
                openListingDialogButton.classList.remove('is-active');
                document.querySelector('.nav-item[data-page="catalog"]').classList.remove('is-active');
                document.querySelector('.nav-item[data-page="account"]').classList.remove('is-active');
                document.querySelector('.page.catalog-page').hidden = false;
                accountPage.hidden = true;

              });
        return;
});