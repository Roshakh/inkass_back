"use strict";


// Dynamic Adapt v.1
// HTML data-da="where(uniq class name),when(breakpoint),position(digi)"
// e.x. data-da=".item,992,2"
// Andrikanych Yevhen 2020
// https://www.youtube.com/c/freelancerlifestyle

function DynamicAdapt(type) {
	this.type = type;
}
DynamicAdapt.prototype.init = function () {
	const _this = this;
	// массив объектов
	this.оbjects = [];
	this.daClassname = "_dynamic_adapt_";
	// массив DOM-элементов
	this.nodes = document.querySelectorAll("[data-da]");

	// наполнение оbjects объктами
	for (let i = 0; i < this.nodes.length; i++) {
		const node = this.nodes[i];
		const data = node.dataset.da.trim();
		const dataArray = data.split(",");
		const оbject = {};
		оbject.element = node;
		оbject.parent = node.parentNode;
		оbject.destination = document.querySelector(dataArray[0].trim());
		оbject.breakpoint = dataArray[1] ? dataArray[1].trim() : "480";
		оbject.place = dataArray[2] ? dataArray[2].trim() : "last";
		оbject.index = this.indexInParent(оbject.parent, оbject.element);
		this.оbjects.push(оbject);
	}
	this.arraySort(this.оbjects);
	// массив уникальных медиа-запросов
	this.mediaQueries = Array.prototype.map.call(this.оbjects, function (item) {
		return '(' + this.type + "-width: " + item.breakpoint + "px)," + item.breakpoint;
	}, this);
	this.mediaQueries = Array.prototype.filter.call(this.mediaQueries, function (item, index, self) {
		return Array.prototype.indexOf.call(self, item) === index;
	});

	// навешивание слушателя на медиа-запрос
	// и вызов обработчика при первом запуске
	for (let i = 0; i < this.mediaQueries.length; i++) {
		const media = this.mediaQueries[i];
		const mediaSplit = String.prototype.split.call(media, ',');
		const matchMedia = window.matchMedia(mediaSplit[0]);
		const mediaBreakpoint = mediaSplit[1];

		// массив объектов с подходящим брейкпоинтом
		const оbjectsFilter = Array.prototype.filter.call(this.оbjects, function (item) {
			return item.breakpoint === mediaBreakpoint;
		});
		matchMedia.addListener(function () {
			_this.mediaHandler(matchMedia, оbjectsFilter);
		});
		this.mediaHandler(matchMedia, оbjectsFilter);
	}
};
DynamicAdapt.prototype.mediaHandler = function (matchMedia, оbjects) {
	if (matchMedia.matches) {
		for (let i = 0; i < оbjects.length; i++) {
			const оbject = оbjects[i];
			оbject.index = this.indexInParent(оbject.parent, оbject.element);
			this.moveTo(оbject.place, оbject.element, оbject.destination);
		}
	} else {
		for (let i = 0; i < оbjects.length; i++) {
			const оbject = оbjects[i];
			if (оbject.element.classList.contains(this.daClassname)) {
				this.moveBack(оbject.parent, оbject.element, оbject.index);
			}
		}
	}
};
// Функция перемещения
DynamicAdapt.prototype.moveTo = function (place, element, destination) {
	element.classList.add(this.daClassname);
	if (place === 'last' || place >= destination.children.length) {
		destination.insertAdjacentElement('beforeend', element);
		return;
	}
	if (place === 'first') {
		destination.insertAdjacentElement('afterbegin', element);
		return;
	}
	destination.children[place].insertAdjacentElement('beforebegin', element);
}
// Функция возврата
DynamicAdapt.prototype.moveBack = function (parent, element, index) {
	element.classList.remove(this.daClassname);
	if (parent.children[index] !== undefined) {
		parent.children[index].insertAdjacentElement('beforebegin', element);
	} else {
		parent.insertAdjacentElement('beforeend', element);
	}
}

// Функция получения индекса внутри родителя
DynamicAdapt.prototype.indexInParent = function (parent, element) {
	const array = Array.prototype.slice.call(parent.children);
	return Array.prototype.indexOf.call(array, element);
};

// Функция сортировки массива по breakpoint и place 
// по возрастанию для this.type = min
// по убыванию для this.type = max
DynamicAdapt.prototype.arraySort = function (arr) {
	if (this.type === "min") {
		Array.prototype.sort.call(arr, function (a, b) {
			if (a.breakpoint === b.breakpoint) {
				if (a.place === b.place) {
					return 0;
				}
				if (a.place === "first" || b.place === "last") {
					return -1;
				}
				if (a.place === "last" || b.place === "first") {
					return 1;
				}
				return a.place - b.place;
			}
			return a.breakpoint - b.breakpoint;
		});
	} else {
		Array.prototype.sort.call(arr, function (a, b) {
			if (a.breakpoint === b.breakpoint) {
				if (a.place === b.place) {
					return 0;
				}
				if (a.place === "first" || b.place === "last") {
					return 1;
				}
				if (a.place === "last" || b.place === "first") {
					return -1;
				}
				return b.place - a.place;
			}
			return b.breakpoint - a.breakpoint;
		});
		return;
	}
};

const da = new DynamicAdapt("max");
da.init();;
$(document).ready(function() {
    if ($('.faq-quest')) {
        $('.item-faqquest__title').next().slideUp(300);
        $('.item-faqquest__title').click(function(event) {
            if($('.faq-quest__list').hasClass('onlyone')) {
                $('.item-faqquest__title').not($(this)).removeClass('active');
                $('.item-faqquest__text').not($(this).next()).slideUp(300);
            }
            $(this).toggleClass('active').next().slideToggle(300);
        })   
    }
    if ($('.contactsp-reg')) {
        $('.item-conreg__title').next().slideUp(300);
        $('.item-conreg__title').click(function(event) {
            if($('.contactsp-reg__list').hasClass('onlyone')) {
                $('.item-conreg__title').not($(this)).removeClass('active');
                $('.item-conreg__body').not($(this).next()).slideUp(300);
            }
            $(this).toggleClass('active').next().slideToggle(300);
        })   
    }
    if ($('.portfolio-quest')) {
        $('.item-portquest__title').next().slideUp(300);
        $('.item-portquest__title').click(function(event) {
            if($('.portfolio__list').hasClass('onlyone')) {
                $('.item-portquest__title').not($(this)).removeClass('active');
                $('.item-portquest__body').not($(this).next()).slideUp(300);
            }
            $(this).toggleClass('active').next().slideToggle(300);
        })   
    }
    if ($('.biography__list.hide')) {
        $('.biography__stitle').prev().slideUp(300);
        $('.biography__stitle').click(function(event) {
            if($('.biography__list').hasClass('onlyone')) {
                $('.biography__stitle').not($(this)).removeClass('active');
                $('.biography__list.hide').not($(this).prev()).slideUp(300);
            }
            $(this).toggleClass('active').prev().slideToggle(300);
        })   
    }
    if ($('.item-newcom')) {
        $('.item-newcom__uncover').prev().slideUp(300);
        $('.item-newcom__uncover').click(function(event) {
            if($('.item-newcom__text').hasClass('onlyone')) {
                $('.item-newcom__uncover').not($(this)).removeClass('active');
                $('.item-newcom__text_hide').not($(this).prev()).slideUp(300);
            }
            $(this).toggleClass('active').prev().slideToggle(300);
        })   
    }
})
;

function ibg(){

let ibg=document.querySelectorAll("._ibg");
    for (var i = 0; i < ibg.length; i++) {
        if(ibg[i].querySelector('img')){
            ibg[i].style.backgroundImage = 'url('+ibg[i].querySelector('img').getAttribute('src')+')';
        }
    }
}

ibg();;
const selectBox = document.querySelectorAll("._select");


selectBox.forEach(element => {

    const selectTag = element.querySelector('select'); 
    const selectTagOptions = selectTag.querySelectorAll('option');
    // const selectTagOptionDisabled = selectTag.querySelector('option[disabled]');

    const iSelect = document.createElement("div");
    const iSelectSelected = document.createElement("span");
    const iSelectList = document.createElement("ul");
    
    iSelect.setAttribute("class", "_select_");
    iSelectSelected.setAttribute("class", "_selected_");
    iSelectList.setAttribute("class", "_select_list");

    element.append(iSelect);
    iSelect.prepend(iSelectSelected);
    iSelect.append(iSelectList);

    selectTagOptions.forEach(option => {
        const iSelectListItem = document.createElement("li");
        iSelectListItem.setAttribute("class", "_select_list_item");
        iSelectList.append(iSelectListItem);
        iSelectListItem.innerHTML = option.innerHTML;
    });

    const iSelectListItems = iSelectList.querySelectorAll('li');

    
    iSelectListItems[0].classList.add('_default');     
    // const iSelectListItemDisabled = iSelectList.querySelector("li._disabled");
    const iSelectListItemDefault = iSelectList.querySelector("li._default");
    
    if(iSelectListItemDefault) {
        iSelectSelected.innerHTML = iSelectListItemDefault.innerHTML;
    }

    let toggleSelect = function() {
        iSelect.classList.toggle('_active');
    }

    let closeSelect = function() {
        iSelect.classList.remove('_active');
    }

    iSelectListItems.forEach(item => {
        item.addEventListener('click', () => {
            // if (iSelectListItemDefault) {
            //     iSelectListItemDefault.remove();   
            // }
            const clickedItemHtml = item.innerHTML;
            iSelectSelected.innerHTML = clickedItemHtml;
            closeSelect();
            if (item.clickedItemHtml !== iSelectSelected) {
                iSelectListItems.forEach(item => {
                    item.classList.remove('hide');
                })
            }
            item.classList.add('hide');
        })
    })

    document.addEventListener('mousedown', function(e){
        if(e.target.closest('._select_') === null){
            closeSelect();
        }
    });

    iSelectSelected.addEventListener('click', toggleSelect);
});





;
let burger = document.querySelector('.nav-header__burger');

let mobileMenu = document.querySelector('.mobile-menu');
let mobileMenuOverlay = document.querySelector('.overlay');
let mobileMenuClose = document.querySelector('.mobile-menu__close');

// let mobileMenuBtn = document.querySelector('.header__burger');
let pageBody = document.querySelector('body');


// let scrollBarWidth = (window.innerWidth - document.body.clientWidth) + 'px';



// console.log(headerBurgerBtn);

burger.addEventListener('click', () => {
    // burger.classList.toggle('_active');

    mobileMenu.classList.toggle('_active');
    mobileMenuOverlay.classList.toggle('_active');

    pageBody.classList.toggle('_lock');
    // document.querySelector('.main').style.paddingRight = scrollBarWidth;
    // document.querySelector('.header').style.paddingRight = scrollBarWidth;
    // document.querySelector('.bottom-header.fix-header').style.paddingRight = scrollBarWidth;
})

mobileMenuOverlay.addEventListener('click', () => {
    // burger.classList.remove('_active');

    mobileMenu.classList.remove('_active');
    mobileMenuOverlay.classList.remove('_active');

    pageBody.classList.remove('_lock');
    // document.querySelector('.main').style.paddingRight = '0';
    // document.querySelector('.header').style.paddingRight = '0';
    // document.querySelector('.bottom-header.fix-header').style.paddingRight = '0';
})


mobileMenuClose.addEventListener('click', () => {
    mobileMenu.classList.remove('_active');
    mobileMenuOverlay.classList.remove('_active');

    pageBody.classList.remove('_lock');
    // document.querySelector('.main').style.paddingRight = '0';
    // document.querySelector('.header').style.paddingRight = '0';
    // document.querySelector('.bottom-header.fix-header').style.paddingRight = '0';
})

;
window.onscroll = () => {
    const headerBottom = document.querySelector('.header__bottom');
    const mobileMenu = document.querySelector('.mobile-menu');
    const mobileMenuClose = document.querySelector('.mobile-menu__close');
    const fixHeaderLogo = document.querySelector('.fix-header__logo');
    const burger = document.querySelector('.nav-header__burger');
    const _scrollY = window.scrollY;
    
    // console.log(_scrollY);
    if(_scrollY > 250) {
        headerBottom.classList.add('fix-header');
    } else if (_scrollY < 250 ) {
        headerBottom.classList.remove('fix-header');
    }
    if (document.querySelector('.first-section')) {
        if (_scrollY > 250) {
            document.querySelector('.header__top').style.paddingTop = '60px';
            // mobileMenu.style.paddingTop = '90px';
            // mobileMenuClose.style.display = 'none';
            // fixHeaderLogo.style.display = 'block';
            // burger.style.marginLeft = '0';
        } else if (_scrollY < 250) {
            document.querySelector('.header__top').style.paddingTop = '0'    
            // mobileMenu.style.paddingTop = '40px';
            // mobileMenuClose.style.top = '32px';
            // fixHeaderLogo.style.display = 'block';
            // burger.style.marginLeft = '50px';
        }
    }
    
};
if (document.querySelector('.swiper')) {
    const swiper1 = new Swiper('.showedswiper',{
        // direction: "vertical",
        loop: true,
        observer: true,
        observeParents: true,
        spaceBetween: 20,
        observeSlideChildren : true,
        navigation: {
            nextEl: ".services__swipermain_nextbtn",
            prevEl: ".services__swipermain_prevbtn",
        },
        clickable: true,    
        breakpoints: {
            0: {
                slidesPerView: 1.1,
            },
            480: {
                spaceBetween: 15,
                slidesPerView: 1.5,
            },
            768: {
                slidesPerView: 3,
            },
            992: {
                slidesPerView: 2,
            }   
        },
    });
    const swiper2 = new Swiper('.controlswiper',{
        // autoHeight: true,
        calculateHeight:true,
        loop: true,
        observer: true,
        observeParents: true,
        observeSlideChildren : true,
        clickable: true,  
        slideToClickedSlide: true,  
        // navigation: {
        //     nextEl: ".",
        //     prevEl: ".",
        // },
        slidesPerView: 3,
        breakpoints: {
            0: {
                direction: "horizontal",
                // mousewheel: false,
                
                spaceBetween: 10,
                slidesPerView: 1.1,
            },
            480: {
                spaceBetween: 15,
                slidesPerView: 1.5,
            },
            768: {
                spaceBetween: 25,
                slidesPerView: 3,
                centeredSlides: false,
            },
            992: {
                centeredSlides: true,
                direction: "vertical",
                // mousewheel: true,
                spaceBetween: 35,
            }
        },
        clickable: true,
    });
    swiper1.controller.control = swiper2;
    swiper2.controller.control = swiper1;
    const swiper3 = new Swiper('.swipernews',{
        // direction: "vertical",
        slidesPerView: 2,
        loop: true,
        observer: true,
        observeParents: true,
        spaceBetween: 20,
        observeSlideChildren : true,
        navigation: {
            nextEl: ".swipernews__nextbtn",
            prevEl: ".swipernews__prevbtn",
        },
        clickable: true,    
        breakpoints: {
            320: {
                slidesPerView: 1.1, 
            },
            480: {
                slidesPerView: 1.4, 
            },
            560: {
                slidesPerView: 2, 
            },
            // 992: {
            //     slidesPerView: 2, 
            // }
        },   
    })
    const swiper4 = new Swiper('.valuta-swiper',{
        loop: true,
        observer: true,
        observeParents: true,
        spaceBetween: 35,
        observeSlideChildren : true,
        clickable: true,
        autoplay: true,
        speed: 2000,
        breakpoints: {
            0: {
                slidesPerView: 1.5, 
            },
            480: {
                slidesPerView: 2, 
            },
            768: {
                slidesPerView: 2.5, 
            },
            992: {
                slidesPerView: 3,
            },
            1270: {
                slidesPerView: 3.4,
            }
        },  
    }) 
    const swiper5 = new Swiper('.comments__swiper',{
        loop: true,
        observer: true,
        observeParents: true,
        spaceBetween: 35,
        observeSlideChildren : true,
        clickable: true, 
        breakpoints: {
            0: {
                slidesPerView: 1, 
            },
            480: {
                slidesPerView: 1.5, 
            },
            640: {
                slidesPerView: 2, 
            },
            768: {
                slidesPerView: 2.5, 
            },
            992: {
                slidesPerView: 3,
            },
        },  
        navigation: {
            nextEl: ".comments__swiper_nextBtn",
        },
        
    })
   
}

;
function hoverImage() {
	document.querySelectorAll('[data-hover-src]').forEach((img) => {
	  const src = img.getAttribute('src');
	  const srcH = img.getAttribute('data-hover-src');
  
	  img.addEventListener('mouseover', () => {img.src = srcH;})
	  img.addEventListener('mouseout', () => {img.src = src;})
	});
}
  
hoverImage();;
// let firtsFullpAgeScroll = new fullpage('#fullpage', {
// 	// menu: '#menu',
// 	// anchors: ['firstPage', 'secondPage', '3rdPage'],
// 	// sectionsColor: ['#C63D0F', '#1BBC9B', '#7E8F7C'],
// 	// autoScrolling: false,
// 	// licenseKey: 'OPEN-SOURCE-GPLV3-LICENSE'
// });;
const tabs = document.querySelectorAll('[data-tab]');
tabs.forEach(element => {
    const btns  = element.querySelectorAll(".manageitem__navbtn");
    const bodys = element.querySelectorAll(".manageitem__tab");
    for (let i = 0; i < btns.length; i++) {
        const btn = btns[i];
        const body = bodys[i];
        btn.addEventListener('click', () => {
            let btnAtribute = btn.getAttribute("data-tab-nav");
            let bodyAtribute = body.getAttribute("data-tab-body");
            if(btnAtribute == bodyAtribute) {
                btns.forEach(btn => {
                    btn.classList.remove('active');
                })
                bodys.forEach(body => {
                    body.classList.remove('active');
                })
                btn.classList.add('active');
                body.classList.add('active');
            }
            console.log(btnAtribute);
            console.log(bodyAtribute);
        })
    }
});
;
let tariftabBodyTr = document.querySelectorAll('.tariftab__body .tr');

function addRemoveActive(onclickBtns) {
    for (let index = 0; index < onclickBtns.length; index++) {
        const element = onclickBtns[index];
        
        element.addEventListener('click', () => {
            onclickBtns.forEach(item => {
                item.classList.remove('_active');
            })
            element.classList.add('_active');
        })
    }
}
addRemoveActive(tariftabBodyTr);
;
$(document).ready(function () {
    let show = true;
    let countbox = $(".whyus-info");
    let numbers = $(".whyus-info__item span");
    if(document.querySelectorAll('.aboutus-section').length) {
        $(window).on("scroll load resize", function () {
            if (!show) return false; // Отменяем показ анимации, если она уже была выполнена
            let w_top = $(window).scrollTop(); // Количество пикселей на которое была прокручена страница
            let e_top = countbox.offset().top; // Расстояние от блока со счетчиками до верха всего документа
            let w_height = $(window).height(); // Высота окна браузера
            let d_height = $(document).height(); // Высота всего документа
            let e_height = countbox.outerHeight(); // Полная высота блока со счетчиками
            if (w_top + 500 >= e_top || w_height + w_top == d_height || e_height + e_top < w_height) {
                numbers.css('opacity', '1');
                numbers.spincrement({
                    thousandSeparator: "",
                    duration: 2500
                });
                show = false;
            }
        });
    }
});






;

const animItems = document.querySelectorAll('._anim-items');

if(animItems.length > 0) {
	window.addEventListener('scroll', animOnScroll)
	function animOnScroll() {
		for (let index = 0; index < animItems.length; index++) {
			const animItem = animItems[index];
			const animItemHeight = animItem.offsetHeight;
			const animItemOffset = offset(animItem).top;
			const animStart = 4;
			const windowWidth = window.innerWidth;

			let animItemPoint  = window.innerHeight - animItemHeight / animStart;

			if (animItemHeight > window.innerHeight) {
				animItemPoint = window.innerHeight - window.innerHeight / animStart;
			}
			if((window.pageYOffset > animItemOffset - animItemPoint) && window.pageYOffset < (animItemOffset + animItemHeight)) {
				if(windowWidth > 768) {
					animItem.classList.add('_active')
				}
			} else {
				if (!animItem.classList.contains('_anim-no-hide')) {
					animItem.classList.remove('_active')	
				}
			}
		}
	}
	function offset (el) {
		const rect = el.getBoundingClientRect(),
		scrollLeft = window.pageXOffset || document.documentElement.scrollLeft,
		scrollTop = window.pageYOffset || document.documentElement.scrollTop;
		return { top: rect.top + scrollTop, left: rect.left + screenLeft}
	}
	// setTimeout(() => {
		
	// }, 500);
	animOnScroll()
}
;



