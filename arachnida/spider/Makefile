# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: minh-ngu <minh-ngu@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2022/09/28 10:56:09 by minh-ngu          #+#    #+#              #
#    Updated: 2024/10/21 17:20:57 by ngoc             ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

MANDA = spider

all:	$(MANDA)
$(MANDA):
	go build -o spider spider.go
example:
	./spider -r -l 3 -p ./images http://example.com
test:
	go test -v
.PHONY: all clean fclean re
