package main

import (
    "fmt"
    "net/http"
)

////////////////////////////////SITI VERI/////////////////////////////////////////////////////////////
func pagemain(w http.ResponseWriter, r *http.Request){
    w.Write([]byte( `
        <!DOCTYPE html>
        <html lang="en">
        <head>
          <meta charset="UTF-8">
          <meta name="viewport" content="width=device-width, initial-scale=1.0">
          <meta http-equiv="X-UA-Compatible" content="ie=edge">
          <!-- cdn vue -->
          <script src="https://cdn.jsdelivr.net/npm/vue@2.6.12/dist/vue.js"></script>
          <!-- cdn axios vue -->
          <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.21.1/axios.min.js" integrity="sha512-bZS47S7sPOxkjU/4Bt0zrhEtWx0y0CRkhEp8IckzK+ltifIIE9EMIMTuT/mEzoIMewUINruDBIR/jJnbguonqQ==" crossorigin="anonymous"></script>
          <!-- cdn fontawesome -->
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css" integrity="sha512-HK5fgLBL+xu6dm/Ii3z4xhlSUyZgTT9tuc/hSrtw6uzJOvgRr2a9jyxxT1ely+B+xFAmJKVSTbpM/CuL7qxO8w==" crossorigin="anonymous" />
          <!-- fonts -->
          <link rel="preconnect" href="https://fonts.gstatic.com">
          <link href="https://fonts.googleapis.com/css2?family=Fredoka+One&family=Source+Sans+Pro:wght@400;700&display=swap" rel="stylesheet">
          <!-- css -->
          <link rel="stylesheet" href="css/style.css">
          <title>Edu</title>
        </head>
        <body>
          <!-- inizio header -->
          <header id="header" :style="{ backgroundImage: 'url(' + bg[indice].url + ')' }">
            <!-- header layer -->
            <div class="layer">
              <!-- inizio menu -->
                <nav class="container">
                  <!-- logo header -->
                  <div class="logo-header">
                    <img src="img/theme_eduprime_logo.png" alt="logo">
                  </div>
                  <!-- fine logo header -->
                  <ul>
                    <!-- lista menu -->
                    <li v-for="(list, index) in menu">
                      <a href="#" @mouseover="showbox(index)">{{list.main}} <i v-if="list.box.length > 1" class="fas fa-sort-down"></i></a>
                      <div v-if="list.box.length > 1" class="menu-box" :class="list.displayBox" @mouseleave="closebox(index)">
                          <a v-for="(item, index) in list.box" href="#">{{item}}</a>
                      </div>
                    </li>
                    <!-- fine lista menu -->
                    <!-- menu button -->
                    <li>
                      <a href="#" class="btn yellow">VIEW COURSES</a>
                    </li>
                    <!-- fine menu button -->
                  </ul>
                  <div class="small-menu">
                    <a href="#"><i class="fas fa-bars"></i></a>
                  </div>
                </nav>
              <!-- fine menu -->
              <!-- inizio jumbotron -->
              <div class="jumbotron container">
                <transition name="slide-fade" mode="out-in">
                  <h1 :key="indice">{{bg[indice].title}}</h1>
                </transition>
                <p>Edu prime is the most versatile WordPress theme for educational purposes,<br> showcasing universities, courses, secondary schools etc. </p>
                <div class="btns">
                  <a href="#" class="btn yellow"><i class="fas fa-search"></i>Search courses</a>
                  <a href="#" class="btn white"><i class="fas fa-user-plus"></i>Apply for university</a>
                </div>
              </div>
              <!-- fine jumbotron -->
              <!-- img bg -->
              <div class="img-bg">
                <img src="img/Wave-1.png" alt="bg img">
              </div>
              <!-- fine img bg -->
              <!-- leftbox -->
              <div class="left-box">
                <a href="#"><i class="fas fa-shopping-cart"></i></a>
                <a href="#"><i class="fas fa-book-open"></i></a>
                <a href="#"><i class="far fa-dot-circle"></i></a>
              </div>
              <!-- fine left box -->
              <!-- inizio corner bottom -->
              <div class="corner-button">
                <a href="#" class="btn yellow"><i class="fas fa-angle-up fa-3x"></i></a>
              </div>
              <!-- fine corner bottom -->
              <!-- inizio arrows -->
              <div class="arrow-left">
                <a href="#" @click="nextSlider"><i class="fas fa-angle-left fa-5x"></i></a>
              </div>
              <div class="arrow-right">
                <a href="#" @click="prevSlider"><i class="fas fa-angle-right fa-5x"></i></a>
              </div>
              <!-- fine arrows -->
            </div>
            <!-- fine header layer -->
          </header>
          <!-- fine header -->
          <!-- inizio main -->
          <main>
            <!-- inizio sezione browse -->
            <section id="browse" >
              <div class="mini-container">
                <div class="bg-img-left">
                  <img src="img/Books-icon.png" alt="bg img">
                </div>
                <div class="content">
                  <div class="text">
                    <h2>Thousends of courses<br> for any type of student</h2>
                    <p>At EduPrime, it doesn’t matter what domain you wish to pursue a career in. Here you can find a course that satisfies your needs within a click away and applies for a course in a matter of minutes. EduPrime is ranked as the most versatile university in the world, thanks to the number of courses it provides.</p>
                    <a href="#" class="btn yellow">Browse through courses</a>
                  </div>
                  <div class="image">
                    <img src="img/Graduation-Illustration.png" alt="graduation img">
                  </div>
                </div>
              </div>
              <div class="bg-img-right">
                <img src="img/home-background.png" alt="bg img">
              </div>
            </section>
            <!-- fine sezione browse -->
            <!-- inizio sezione faculties -->
            <div id="faculties">
              <h2>Faculties available at EduPrime</h2>
              <p>A single university with a load of courses, tailored<br> to satisfy any student's needs</p>
                <div class="images">
                  <div v-for="(item, index) in faculties" class="img" :class="item.display" @click="showContent(index)">
                    <img :src="item.smallImg" alt="faculties">
                    <h3>{{item.name}}</h3>
                    <i v-if="item.display== 'active'" class="fas fa-sort-down fa-3x"></i>
                  </div>
                </div>
                <div class="content mini-container">
                  <div class="image">
                    <img :src="faculties[indice].bigImg" alt="faculties">
                  </div>
                  <div class="text">
                    <h2>{{faculties[indice].name}}</h2>
                    <p>{{faculties[indice].text}}</p>
                    <a href="#" class="btn yellow">Read More</a>
                  </div>
                </div>
                <div class="img-bg">
                  <img src="svg/svg-0.svg" alt="bg">
                </div>
            </div>
            <!-- fine sezione faculties -->
            <!-- inizio sezione year -->
            <div id="year">
                <div class="clock-img">
                  <img src="img/Clock-and-Bell.png" alt="clock">
                </div>
                <h2>University Year</h2>
                <!-- first row -->
                <div class="row first">
                  <div class="text">
                    <h3>Demo Classes</h3>
                    <p>In the first week, students try to accommodate with the teaching style and choose their optional courses.</p>
                  </div>
                  <div class="text">
                    <h3>Graduation Day</h3>
                    <p>On the day of graduation, all students gather for the ceremony and then network and party among others.</p>
                  </div>
                </div>
                <!-- fine primo row -->
                <!-- timeline -->
                <div class="timeline">
                  <img src="img/Timeline-Item.png" alt="timeline">
                </div>
                <!-- fine timeline -->
                <!-- secondo row -->
                <div class="row last">
                  <div class="text">
                    <h3>Orientation</h3>
                    <p>First day of the university year, all students gather for the opening ceremony and then network with others.</p>
                  </div>
                  <div class="text">
                    <h3>Evaluation</h3>
                    <p>At the end of a semester, students take a general evaluation test for every subject they’re learning.</p>
                  </div>
                </div>
                <!-- fine secondo row -->
            </div>
            <!-- fine sezione year -->
            <!-- inizio sezione events -->
            <div id="events">
              <div class="calendar-img">
                <img src="img/upcoming-events-calendar-icon.png" alt="calendar">
              </div>
              <h2>Upcoming Events</h2>
              <div class="row">
                <div class="event">
                  <h3>Coaching Sessions</h3>
                  <p><i class="far fa-calendar-alt"> 20 May 21:30 PM</i></p>
                  <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                  <a href="#" class="btn red"><i class="fas fa-plus"></i> Find More</a>
                </div>
                <div class="event">
                  <h3>Coaching Sessions</h3>
                  <p><i class="far fa-calendar-alt"> 24 Mar 18:00 PM</i></p>
                  <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                  <a href="#" class="btn red"><i class="fas fa-plus"></i> Find More</a>
                </div>
                <div class="event">
                  <h3>Coaching Sessions</h3>
                  <p><i class="far fa-calendar-alt"> 12 Feb 13:30 PM</i></p>
                  <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                  <a href="#" class="btn red"><i class="fas fa-plus"></i> Find More</a>
                </div>
              </div>
              <a href="#" class="btn yellow">View All Events</a>
            </div>
            <!-- fine sezione events -->
            <!-- inizio sezione latest -->
            <div id="latest">
              <div class="mini-container">
                <h2>Latest Courses</h2>
                <div class="row">
                  <div class="course">
                    <div class="course-img">
                      <img src="img/Decisions-icon.png" alt="img">
                    </div>
                    <h3>Make Better Decisions</h3>
                    <p><i class="far fa-user"></i>Teacher: <span>James Collins</span></p>
                    <p><i class="fas fa-money-bill"></i>Price: <span>Free</span></p>
                    <a href="#" class="btn red"><i class="fas fa-eye"></i> View Course</a>
                  </div>
                  <div class="course">
                    <div class="course-img">
                      <img src="img/Decisions-icon.png" alt="img">
                    </div>
                    <h3>How to be a speaker</h3>
                    <p><i class="far fa-user"></i>Teacher: <span>James Collins</span></p>
                    <p><i class="fas fa-money-bill"></i>Price: <span>Free</span></p>
                    <a href="#" class="btn red"><i class="fas fa-eye"></i> View Course</a>
                  </div>
                  <div class="course">
                    <div class="course-img">
                      <img src="img/Decisions-icon.png" alt="img">
                    </div>
                    <h3>Network Introductions</h3>
                    <p><i class="far fa-user"></i>Teacher: <span>James Collins</span></p>
                    <p><i class="fas fa-money-bill"></i>Price: <span>Free</span></p>
                    <a href="#" class="btn red"><i class="fas fa-eye"></i> View Course</a>
                  </div>
                  <div class="course">
                    <div class="course-img">
                      <img src="img/Decisions-icon.png" alt="img">
                    </div>
                    <h3>Brand Management</h3>
                    <p><i class="far fa-user"></i>Teacher: <span>James Collins</span></p>
                    <p><i class="fas fa-money-bill"></i>Price: <span>Free</span></p>
                    <a href="#" class="btn red"><i class="fas fa-eye"></i> View Course</a>
                  </div>
                  <div class="course">
                    <div class="course-img">
                      <img src="img/Decisions-icon.png" alt="img">
                    </div>
                    <h3>Make Better Decisions</h3>
                    <p><i class="far fa-user"></i>Teacher: <span>James Collins</span></p>
                    <p><i class="fas fa-money-bill"></i>Price: <span>Free</span></p>
                    <a href="#" class="btn red"><i class="fas fa-eye"></i> View Course</a>
                  </div>
                </div>
              </div>
              <div class="img-bg">
                <img src="svg/svg-1.svg" alt="bg">
              </div>
            </div>
            <!-- fine sezione latest -->
            <!-- inizio sezione method -->
            <section id="method" >
              <!-- first container -->
              <div class="mini-container">
                <div class="bg-img-left">
                  <img src="img/Exam-icon.png" alt="bg img">
                </div>
                <div class="content">
                  <div class="text">
                    <h2>The most efficient<br> examination method</h2>
                    <p>EduPrime has gathered teachers from around the globe to brainstorm in order to facilitate the evaluation of our students. Every teacher from our university has an influence on how students are evaluated at his/her subject.</p>
                    <a href="#" class="btn red">Discover the Method</a>
                  </div>
                  <div class="image">
                    <img src="img/Exam-Illustration.png" alt="graduation img">
                  </div>
                </div>
              </div>
              <!-- fine first container -->
              <!-- second container -->
              <div class="mini-container">
                <div class="bg-img-right">
                  <img src="img/Exam-icon-1.png" alt="bg img">
                </div>
                <div class="content">
                  <div class="image">
                    <img src="img/Girl-Illustration.png" alt="graduation img">
                  </div>
                  <div class="text">
                    <h2>Variable fees for<br> international students</h2>
                    <p>EduPrime has gathered teachers from around the globe to brainstorm in order to facilitate the evaluation of our students. Every teacher from our university has an influence on how students are evaluated at his/her subject.</p>
                    <a href="#" class="btn red">List of fees</a>
                  </div>
                </div>
              </div>
              <!-- fine second container -->
            </section>
            <!-- fine sezione method -->
            <!-- inizio sezione newsletter -->
            <div id="newsletter">
              <h2>Subscribe Now to Our Newsletter</h2>
              <div class="email">
                <input type="text" placeholder="Email address...">
                <a href="#" class="btn yellow"><i class="fas fa-paper-plane fa-lg"></i></a>
              </div>
            </div>
            <!-- fine sezione newsletter -->
            <!-- inizio sezione partners -->
            <div id="partners">
              <h2>Partners</h2>
              <p>Leverage agile frameworks to provide a robust synopsis for high level<br> overviews. Iterative approaches to corporate strategy.</p>
              <div class="images mini-container">
                <img src="img/partner-1.png" alt="partners">
                <img src="img/partner-2.png" alt="partners">
                <img src="img/partner-3.png" alt="partners">
                <img src="img/partner-4.png" alt="partners">
                <img src="img/partner-5.png" alt="partners">
                <img src="img/partner-6.png" alt="partners">
                <img src="img/partner-7.png" alt="partners">
                <img src="img/partner-8.png" alt="partners">
              </div>
              <div class="bg-img">
                <img src="img/background-wave3.png" alt="">
              </div>
            </div>
            <!-- fine sezione partners -->
          </main>
          <!-- fine main -->
          <!-- inizio footer -->
          <footer>
            <div class="mini-container">
              <div class="footer-left">
                <div class="logo">
                  <img src="img/theme_eduprime_logo.png" alt="logo">
                </div>
                <p>EduPrime is the most versatile WordPress theme for<br> educational purposes, showcasing universities,<br> courses, secondary schools etc.</p>
                <div class="social">
                  <a href="#" class="btn red"><i class="fab fa-facebook-f"></i></a>
                  <a href="#" class="btn red"><i class="fab fa-twitter"></i></a>
                  <a href="#" class="btn red"><i class="fab fa-instagram"></i></a>
                </div>
              </div>
              <div class="footer-middle">
                <h3>Get EduPrime</h3>
                <a href="#"><p>Request a website</p></a>
                <a href="#"><p>Browse Themes</p></a>
                <a href="#"><p>Payment options</p></a>
                <a href="#"><p>Support System</p></a>
                <a href="#"><p>Checkout</p></a>
                <a href="#"><p>Purchase Theme</p></a>
              </div>
              <div class="footer-middle">
                <h3>Networking</h3>
                <a href="#"><p>Purchase Theme</p></a>
                <a href="#"><p>Our Benefits</p></a>
                <a href="#"><p>Our Team</p></a>
                <a href="#"><p>Our Services</p></a>
                <a href="#"><p>Other Products</p></a>
                <a href="#"><p>My account</p></a>
              </div>
              <div class="footer-right">
                <div class="searchbar">
                  <input type="text" placeholder="Search...">
                  <a href="#" class="btn yellow"><i class="fas fa-search"></i></a>
                </div>
                <h3>Search categories</h3>
                <p>
                  <a href="#"><span>ECONOMY</span></a>
                  <a href="#"><span>DESIGN</span></a>
                  <a href="#"><span>COACHING</span></a>
                  <a href="#"><span>BUSINESS</span></a>
                </p>
                <p>
                  <a href="#"><span>MEDICINE</span></a>
                  <a href="#"><span>LAW</span></a>
                  <a href="#"><span>FITNESS</span></a>
                </p>
                <p>ModelTheme. All rights reserved.</p>
              </div>
            </div>
          </footer>
          <!-- fine footer -->
        <!-- javascript -->
        <script src="js/script.js" charset="utf-8"></script>
        </body>
        </html>
        `))
    }

////////////////////////////////SITI VERI/////////////////////////  ////////////////////////////////////

            func main(){
                http.HandleFunc("/",pagemain)

                fmt.Println("Listening on http://localhost:3000/")
                    http.ListenAndServe(":3000",nil)

            }
