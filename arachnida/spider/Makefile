# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: minh-ngu <minh-ngu@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2022/09/28 10:56:09 by minh-ngu          #+#    #+#              #
#    Updated: 2024/11/04 10:34:51 by ngoc             ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

SRCS = *.go
MANDA = spider
DATA_DIR := ./data

all:	$(MANDA)
$(MANDA):
	make build
clean:
	@echo "Cleaning $(DATA_DIR) directory..."
	@rm -rf $(DATA_DIR)/*
	@echo "Cleaned $(DATA_DIR) directory."
build:
	@echo "Building spider program..."
	@go build -o spider spider.go
	@echo "Built spider program."
test:
	@echo "Running tests..."
	@go test -v
	@echo "Tests completed."
run:
	@echo "Running spider program..."
	@./spider -r -l 3 -p ./data http://example.com
	@echo "Spider program completed."
.PHONY: all clean fclean re
