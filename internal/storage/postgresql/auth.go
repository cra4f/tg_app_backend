package postgresql

import (
	"fmt"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (p *Postgresql) Login(user *initdata.User) error {
	fmt.Println("user", user)

	request := fmt.Sprintf(
		`DO
				$do$
				BEGIN
				   IF NOT EXISTS (SELECT FROM public.users WHERE id = %d) THEN
					 INSERT INTO public.users(id, first_name, last_name, username, lang_code, added_to_attachment_menu, allows_write_to_pm, is_bot, is_premium, photo_url) 
					 VALUES (%d, '%s', '%s', '%s', '%s', %t, %t, %t, %t, '%s');
				   END IF;
				END
				$do$`,
		user.ID,
		user.ID, user.FirstName, user.LastName, user.Username,
		user.LanguageCode, user.AddedToAttachmentMenu, user.AllowsWriteToPm,
		user.IsBot, user.IsPremium, user.PhotoURL)

	stmt, err := p.db.Prepare(request)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}
