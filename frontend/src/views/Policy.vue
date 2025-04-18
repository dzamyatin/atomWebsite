<template>
    <div>
      <section class="section-hero section-shaped my-0">
        <div class="shape shape-style-1 shape-primary">
          <span class="span-150"></span>
          <span class="span-50"></span>
          <span class="span-50"></span>
          <span class="span-75"></span>
          <span class="span-100"></span>
          <span class="span-75"></span>
          <span class="span-50"></span>
          <span class="span-100"></span>
          <span class="span-50"></span>
          <span class="span-100"></span>
        </div>
        <div class="container shape-container d-flex align-items-center">
          <div class="col px-0">
            <div class="row align-items-center justify-content-around stars-and-coded">
              <div class="col-sm-4">
                <span class="text-white alpha-7 ml-3">Star us on</span>
                <a href="https://github.com/creativetimofficial/argon-design-system" target="_blank" title="Support us on Github">
                  <img src="/img/brand/github-white-slim.png" style="height: 22px; margin-top: -3px">
                </a>
              </div>
              <div class="col-sm-4 mt-4 mt-sm-0 text-right">
                <span class="text-white alpha-7">Coded by</span>
                <a href="https://www.creative-tim.com" target="_blank" title="Creative Tim - Premium Bootstrap Themes and Templates">
                  <img src="/img/brand/creativetim-white-slim.png" class="ml-3" style="height: 30px;">
                </a>
              </div>
            </div>
          </div>
        </div>
      </section>
        <div class="profile-page">
          <section class="section section-skew">
            <div class="container">
              <card shadow class="card-profile" no-body>
                <div class="px-4">
                  <div class="text-center mt-5">
                    <h3>
                      Политика
                    </h3>
                  </div>
                  <div class="mt-5 py-5 border-top ">
                    <div class="row justify-content-center">
                      <div class="col-lg-9">
                        <pre>
Политика безопасности ВПН-провайдера ЗАО «Итсервис»
1. Введение
Политика безопасности ЗАО «Итсервис» направлена на обеспечение конфиденциальности, целостности и доступности данных пользователей, а также на соблюдение всех требований российского законодательства.

2. Обязательства по конфиденциальности

Безлоговая политика: ЗАО «Итсервис» придерживается политики отсутствия логов, что означает, что мы не храним логи активности пользователей. Это гарантирует, что данные пользователей не могут быть переданы третьим лицам или использованы в коммерческих целях.
Защита данных: Все данные, передаваемые через наши серверы, шифруются с использованием современных протоколов шифрования, таких как OpenVPN, что обеспечивает высокий уровень безопасности.
3. Соблюдение законодательства

Сотрудничество с госорганами: В случае законных запросов от российских государственных органов, ЗАО «Итсервис» обязуется предоставлять необходимую информацию в соответствии с действующим законодательством.
Прозрачность: Пользователи будут уведомлены о любых изменениях в политике безопасности или случаях, когда данные могут быть переданы по запросу госорганов.
4. Технические меры безопасности

Шифрование трафика: Весь интернет-трафик пользователей шифруется, что предотвращает перехват данных злоумышленниками.
Защита от утечек DNS: Мы используем собственные DNS-серверы для предотвращения утечек DNS-запросов, что дополнительно защищает конфиденциальность пользователей.
Регулярные аудиты: Проводим регулярные аудиты безопасности для выявления и устранения уязвимостей.
5. Поддержка пользователей

Обратная связь: Пользователи могут в любое время связаться с нашей службой поддержки для получения информации о мерах безопасности и задать вопросы.
Обучение: Предоставляем пользователям рекомендации по настройке ВПН и обеспечению безопасности их устройств.
6. Заключение

ЗАО «Итсервис» стремится обеспечить максимальную безопасность и конфиденциальность для всех своих пользователей, соблюдая при этом все требования российского законодательства.
                        </pre>
                      </div>
                    </div>
                  </div>
                </div>
              </card>
            </div>
          </section>
        </div>
    </div>
</template>
<script>
import MainTop from "./components/MainTop.vue";
import MainBottom from "./components/MainBottom.vue";
// import Hero from "./components/Hero";

export default {
  name: "components",
  components: {
    MainTop,
    MainBottom,
    // Hero,
  }
};
</script>
